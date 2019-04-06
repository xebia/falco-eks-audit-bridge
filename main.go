package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/jpillora/backoff"
)

const (
	// Use the eu-west-1 region as default
	awsDefaultRegion = "eu-west-1"

	// Firehose will create objects with a year/month/day/hour structure.
	// Prefix indicates that we are only interested in the events from
	// firehose and not anything else in the bucket.
	firehosePrefix = "20"

	// dataEventMessageType is the Firehose message type for data
	dataEventMessageType = "DATA_MESSAGE"

	// checkInterval is the delay in minutes to check for new Firehose events
	checkInterval = 2 * time.Minute
)

// Log is a CloudWatch K8s audit log line
type Log struct {
	Message string `json:"message"`
}

// Event is a firehose event
type Event struct {
	MessageType string `json:"messageType"`
	LogEvents   []Log  `json:"logEvents"`
}

func main() {

	// Initialization

	bucket, ok := os.LookupEnv("BUCKET")
	if !ok {
		fmt.Println("Environment variable 'BUCKET' not set, exiting.")
		os.Exit(1)
	}

	falcoEndpoint, ok := os.LookupEnv("FALCO_ENDPOINT")
	if !ok {
		fmt.Println("Environment variable 'FALCO_ENDPOINT' not set, exiting.")
	}

	region, ok := os.LookupEnv("AWS_DEFAULT_REGION")
	if !ok {
		region = awsDefaultRegion
	}

	prefix, ok := os.LookupEnv("FIREHOSE_PREFIX")
	if !ok {
		prefix = firehosePrefix
	}

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))
	s3client := s3.New(sess)

	b := &backoff.Backoff{
		Jitter: true,
	}

	httpClient := retryablehttp.NewClient()

	// Disable the verbose debug logging for now
	httpClient.Logger = nil

	// Start of Firehove events processing

	for {
		res, err := s3client.ListObjects(&s3.ListObjectsInput{
			Bucket: aws.String(bucket),
			Prefix: aws.String(prefix),
		})

		if err != nil {
			d := b.Duration()
			fmt.Printf("Error listing bucket:\n%v\nRetrying in %s", err, d)
			time.Sleep(d)
			continue
		}
		b.Reset()

		objects := res.Contents

		// Sort the objects according to LastModified date
		sort.Slice(objects, func(i, j int) bool {
			return objects[i].LastModified.Before(*objects[j].LastModified)
		})

		for _, object := range objects {

			// Check if the object was already processed but somehow not deleted
			_, err := s3client.HeadObject(&s3.HeadObjectInput{
				Bucket: aws.String(bucket),
				Key:    aws.String(fmt.Sprintf("processed/%s", *object.Key)),
			})

			if err != nil {
				// Apparently the object doesn't exist in the processed folder, so start processing it
				fmt.Printf("Processing: %s\n", *object.Key)
			} else {
				// The object already exists in the processed folder, so delete and ignore it
				_, err = s3client.DeleteObject(&s3.DeleteObjectInput{
					Bucket: aws.String(bucket),
					Key:    object.Key,
				})

				if err != nil {
					fmt.Printf("Could not delete file from the Firehose folder:\n%v\n", err)
					continue
				}
			}

			// Download the file from S3
			file, err := s3client.GetObject(
				&s3.GetObjectInput{
					Bucket: aws.String(bucket),
					Key:    object.Key,
				})
			if err != nil {
				fmt.Printf("Could not download the Firehose events file:\n%v\n", err)
				continue
			}

			// Uncompress the file
			contents, err := gzip.NewReader(file.Body)
			if err != nil {
				fmt.Printf("Could not decompress the Firehose events:\n%v\n", err)
				continue
			}

			// Parse the JSON file contents
			decoder := json.NewDecoder(contents)
			event := Event{}

			var processed bool
		DECODER:
			for {

				if err := decoder.Decode(&event); err == io.EOF {
					processed = true
					break
				} else if err != nil {
					fmt.Printf("Unable to (fully) parse Firehose event:\n%v\n", err)
					break
				}

				if event.MessageType == dataEventMessageType {
					for _, log := range event.LogEvents {

						// Post the audit log to Falco for compliance checking
						res, err := httpClient.Post(falcoEndpoint, "application/json", strings.NewReader(log.Message))
						if err != nil || res.StatusCode != 200 {
							fmt.Printf("Unable to send the audit log to Falco:\n%v\n", err)
							break DECODER
						}
						res.Body.Close()
					}
				}
			}

			if processed {
				_, err = s3client.CopyObject(&s3.CopyObjectInput{
					Bucket:     aws.String(bucket),
					CopySource: aws.String(fmt.Sprintf("/%s/%s", bucket, *object.Key)),
					Key:        aws.String(fmt.Sprintf("processed/%s", *object.Key)),
				})

				if err != nil {
					fmt.Printf("Could not copy file to processed folder:\n%v\n", err)
					continue
				}

				_, err = s3client.DeleteObject(&s3.DeleteObjectInput{
					Bucket: aws.String(bucket),
					Key:    object.Key,
				})

				if err != nil {
					fmt.Printf("Could not delete file from the Firehose folder:\n%v\n", err)
					continue
				}
			} else {
				fmt.Printf("Object '%s' was not (fully) processed, not moving it to the processed folder.\n", *object.Key)
			}
		}

		if len(objects) == 0 {
			fmt.Println("No new Firehose events found, waiting for next check interval.")
			time.Sleep(checkInterval)
		}
	}
}
