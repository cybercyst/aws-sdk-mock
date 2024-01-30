package main

import (
	"io"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// localstack has a lot of logic to detect which service to use
		// but i believe for our services we can parse authorization
		// content= for the service, which is the 4th chunk
		authorization := r.Header.Get("authorization")

		switch {
		case strings.Contains(authorization, "s3"):
			_, _ = io.WriteString(w, MockListBuckets)
		case strings.Contains(authorization, "ec2"):
			_, _ = io.WriteString(w, MockDescribeInstances)
		}
	})

	err := http.ListenAndServe(":4566", nil)
	if err != nil {
		panic(err)
	}
}
