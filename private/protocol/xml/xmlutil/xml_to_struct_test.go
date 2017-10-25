package xmlutil_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awstesting/unit"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func TestUnmarshal(t *testing.T) {
	xmlVal := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<AccessControlPolicy xmlns="http://s3.amazonaws.com/doc/2006-03-01/">
	<Owner>
		<ID>foo-id</ID>
		<DisplayName>user</DisplayName>
	</Owner>
	<AccessControlList>
		<Grant>
		<Grantee xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:type="type">
			<ID>foo-id</ID>
			<DisplayName>user</DisplayName>
		</Grantee>
		<Permission>FULL_CONTROL</Permission>
		</Grant>
	</AccessControlList>
</AccessControlPolicy>`)

	var server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(xmlVal)
	}))

	cfg := unit.Config()
	cfg.EndpointResolver = aws.ResolveWithEndpointURL(server.URL)
	cfg.S3ForcePathStyle = aws.Bool(true)

	svc := s3.New(cfg)

	out, err := svc.GetBucketAcl(&s3.GetBucketAclInput{
		Bucket: aws.String("foo"),
	})

	assert.NoError(t, err)

	expected := &s3.GetBucketAclOutput{
		Grants: []*s3.Grant{
			{
				Grantee: &s3.Grantee{
					DisplayName: aws.String("user"),
					ID:          aws.String("foo-id"),
					Type:        aws.String("type"),
				},
				Permission: aws.String("FULL_CONTROL"),
			},
		},

		Owner: &s3.Owner{
			DisplayName: aws.String("user"),
			ID:          aws.String("foo-id"),
		},
	}
	assert.Equal(t, expected, out)
}
