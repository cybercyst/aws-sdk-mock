This repo serves to demonstrate discovering and mocking an AWS API

`aws-app` contains an app using the `S3` SDK in order to ListBuckets you can see

```
$ cd aws-app

# Authenticate as a user with some S3 buckets in AWS

$ go run .

bucket1
bucket2
...
```

Let's say we'd like to create an mock API server for integration tests...

```
$ cd aws-mock

$ go run .

# in another terminal, split plane, etc

$ cd aws-app

$ AWS_ACCESS_KEY_ID=test AWS_SECRET_ACCESS_KEY=test AWS_REGION=us-east-1 AWS_ENDPOINT=http://localhost:4566 go run .
mock-bucket-1
```
