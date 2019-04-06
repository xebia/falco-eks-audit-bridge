# Falco EKS audit bridge

This tool reads an S3 bucket which contains AWS CloudWatch audit logs coming from EKS. The files in S3 are coming from AWS Kinesis, see [here](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs//SubscriptionFilters.html#FirehoseExample) for more info how to set that up.

## Build

```bash
go build
```