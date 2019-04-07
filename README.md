# Falco EKS Audit Bridge

Falco EKS Audit Bridge (FEAB) monitors an S3 bucket with AWS EKS audit logs and sends them to Sysdig Falco for inspection. The EKS audit logs are retrieved from AWS CloudWatch with the AWS Kinesis Firehose service. Please check this [guide](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs//SubscriptionFilters.html#FirehoseExample) to learn how to setup the AWS infrastructure.

## Building

In order to build this project you can use the Makefile, which contains two goals:

1. **bin**: create a local, system specific, binary with the local Go installation (1.11+ required)
2. **docker**: create a docker image from scratch with the tool added

## Configuration

Two environment variables are required when starting the bridge:

1. **BUCKET**: the S3 bucket to monitor. Please make sure that the correct AWS credentials are available to the application (either as environment variables, instance profile, etc.)
2. **FALCO_ENDPOINT**: the Falco HTTP endpoint (e.g. http://localhost:8765/k8s_audit)

## Deployment

This tool is meant to run as a service within Kubernetes. To that end, we have provided a Helm chart which makes deployment easy. You can of course create your own deployment configuration for any system with the docker image.

The Helm chart contains several configuration options that you can override for your specific environment.