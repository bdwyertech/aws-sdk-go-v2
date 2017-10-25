// +build go1.8

package corehandlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/defaults"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func TestSendHandler_HEADNoBody(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	cfg := defaults.Config()
	cfg.Region = aws.String("mock-region")
	cfg.EndpointResolver = aws.ResolveWithEndpointURL(server.URL)
	cfg.DisableSSL = aws.Bool(true)
	cfg.S3ForcePathStyle = aws.Bool(true)

	svc := s3.New(cfg)

	req, _ := svc.HeadObjectRequest(&s3.HeadObjectInput{
		Bucket: aws.String("bucketname"),
		Key:    aws.String("keyname"),
	})

	if e, a := aws.NoBody, req.HTTPRequest.Body; e != a {
		t.Fatalf("expect %T request body, got %T", e, a)
	}

	err := req.Send()
	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}
	if e, a := http.StatusOK, req.HTTPResponse.StatusCode; e != a {
		t.Errorf("expect %d status code, got %d", e, a)
	}
}
