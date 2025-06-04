package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
)

func uploadFile(client *s3.Client) {
	bucket := os.Getenv("S3_BUCKET")
	key := "output.pdf"
	file, err := os.Open("C:\\dev\\output.pdf")
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   file,
	})
	if err != nil {
		log.Fatalf("failed to upload file: %v", err)
	}
	fmt.Printf("File %s uploaded to bucket %s\n", key, bucket)
}

func downloadFile(client *s3.Client) {
	bucket := os.Getenv("S3_BUCKET")
	key := "output.pdf"

	output, err := client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		log.Fatalf("failed to download file: %v", err)
	}
	defer output.Body.Close()

	outFile, err := os.Create("C:\\dev\\downloaded_output.pdf")
	if err != nil {
		log.Fatalf("failed to create file: %v", err)
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, output.Body)
	if err != nil {
		log.Fatalf("failed to save file: %v", err)
	}
	fmt.Printf("File %s downloaded from bucket %s\n", key, bucket)
}

func listBuckets(client *s3.Client) {
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
}

func listObjects(client *s3.Client) {
	// Example: List objects in a bucket (replace "your-bucket-name")
	bucket := os.Getenv("S3_BUCKET")
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

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Load AWS config from environment or shared config
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-2"),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				os.Getenv("YOUR_ACCESS_KEY_ID"),
				os.Getenv("YOUR_SECRET_ACCESS_KEY"),
				"",
			),
		),
	)
	if err != nil {
		log.Fatalf("unable to conenct to aws, %v", err)
	}

	// Create S3 client
	client := s3.NewFromConfig(cfg)

	// listBuckets(client)
	// listObjects(client)
	// uploadFile(client)
	downloadFile(client)
}
