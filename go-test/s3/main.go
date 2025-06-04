package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	// Load AWS config from environment or shared config
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-2"),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				"YOUR_ACCESS_KEY_ID",
				"YOUR_SECRET_ACCESS_KEY",
				"",
			),
		),
	)
	if err != nil {
		log.Fatalf("unable to conenct to aws, %v", err)
	}

	// Create S3 client
	client := s3.NewFromConfig(cfg)

	// Example: List buckets
	// result, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	result, err := client.ListDirectoryBuckets(context.TODO(), nil)
	if err != nil {
		log.Fatalf("unable to list buckets, %v", err)
	}

	fmt.Println("Buckets:")
	for _, b := range result.Buckets {
		fmt.Printf("* %s\n", aws.ToString(b.Name))
	}

	// Example: List objects in a bucket (replace "your-bucket-name")
	bucket := "your-bucket-name"
	objs, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		log.Fatalf("unable to list objects, %v", err)
	}
	fmt.Printf("Objects in bucket %s:\n", bucket)
	for _, obj := range objs.Contents {
		fmt.Println(*obj.Key)
	}
}
