package main

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func newS3Client(accessKey, secret string) *s3.Client {
	cfg := aws.Config{
		Region:      "us-east-1",
		Credentials: credentials.NewStaticCredentialsProvider(accessKey, secret, ""),
	}
	return s3.NewFromConfig(cfg)
}
