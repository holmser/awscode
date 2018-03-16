# AWSCode

This project aims to be a simple cli tool using [Cobra](https://github.com/spf13/cobra) and the [AWS Go SDK](https://github.com/aws/aws-sdk-go) to interface with AWS CodeCommit, CodeBuild, CodePipeline, and CodeDeploy.  

### Usage 
Requirements:
- AWS credentials set up according to [this doc](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-getting-started.html)
- Go Environment set up correctly


```
# Download the package
go get github.com/holmser/awscode

# Install the binary
go install github.com/holmser/awscode

# Get help
awscode help
```
