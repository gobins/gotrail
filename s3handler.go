package gotrail

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func getBuckets() []*s3.Bucket {
	s3client := gets3client()
	var params *s3.ListBucketsInput
	resp, err := s3client.ListBuckets(params)

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(resp)
	return resp.Buckets
}

func getObjects() {
	s3client := gets3client()
	params := &s3.ListObjectsInput{
		Bucket: aws.String("BucketName"), // Required
		// Delimiter:    aws.String("Delimiter"),
		// EncodingType: aws.String("EncodingType"),
		// Marker:       aws.String("Marker"),
		// MaxKeys:      aws.Int64(1),
		// Prefix:       aws.String("Prefix"),
	}
	resp, err := s3client.ListObjects(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
