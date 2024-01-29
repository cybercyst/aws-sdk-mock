package main

import (
	"encoding/xml"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type ListAllMyBucketsResult struct {
	s3.ListBucketsOutput
}

var (
	MockListBucketsResponse = ListAllMyBucketsResult{
		ListBucketsOutput: s3.ListBucketsOutput{
			Buckets: []types.Bucket{
				{
					CreationDate: aws.Time(time.Now()),
					Name:         aws.String("mock-bucket"),
				},
			},
			Owner: &types.Owner{
				ID:          aws.String("AIDACKCEVSQ6C2EXAMPLE"),
				DisplayName: aws.String("Account+Name"),
			},
		},
	}
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, err := xml.Marshal(MockListBucketsResponse)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Internal Server Error"))
		}

		_, _ = w.Write(data)
		// 		_, _ = io.WriteString(w, `<?xml version='1.0' encoding='utf-8'?>
		// <ListAllMyBucketsResult xmlns="http://s3.amazonaws.com/doc/2006-03-01/"><Owner><DisplayName>webfile</DisplayName><ID>75aa57f09aa0c8caeab4f8c24e99d10f8e7faeebf76c078efc7c6caea54ba06a</ID></Owner><Buckets><Bucket><Name>test</Name><CreationDate>2024-01-29T11:33:51.000Z</CreationDate></Bucket></Buckets></ListAllMyBucketsResult>
		// `)
	})

	err := http.ListenAndServe(":4566", nil)
	if err != nil {
		panic(err)
	}
}
