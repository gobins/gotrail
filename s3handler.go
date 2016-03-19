package gotrail

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/s3"
)

func getBuckets() []*s3.Bucket {
	s3client := gets3client()
	var params *s3.ListBucketsInput
	resp, err := s3client.ListBuckets(params)

	if err != nil {
		fmt.Println(err.Error())
	}
	return resp.Buckets
}
