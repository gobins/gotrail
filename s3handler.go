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
	params := &s3.GetObjectInput{
		Bucket: aws.String("BucketName"),
		Key:    aws.String("ObjectKey"),
	}
	resp, err := s3client.GetObject(params)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func downloadObjects() {

}
