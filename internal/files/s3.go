package files

import (
	"bytes"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/jonnypillar/somniloquy/configs"
	"github.com/pkg/errors"
)

// Bucket ...
type Bucket struct {
	bucket     *s3.S3
	bucketName string
}

// NewBucket ...
func NewBucket(config *config.ServiceConfig) (*Bucket, error) {
	creds := credentials.NewCredentials(&credentials.SharedCredentialsProvider{})
	_, err := creds.Get()
	if err != nil {
		return nil, errors.Wrap(err, "error occured creating AWS credentials")
	}
	cfg := aws.NewConfig().WithRegion(config.AWSRegion).WithCredentials(creds)
	svc := s3.New(session.New(), cfg)

	return &Bucket{
		bucket:     svc,
		bucketName: config.AWSUploadS3BucketName,
	}, nil
}

// Upload ...
func (b Bucket) Upload(filename string, by *bytes.Buffer) error {
	data := by.Bytes()

	fileBytes := bytes.NewReader(data)
	path := "/media/" + filename

	params := &s3.PutObjectInput{
		Bucket:        aws.String(b.bucketName),
		Key:           aws.String(path),
		Body:          fileBytes,
		ContentLength: aws.Int64(int64(len(data))),
		ContentType:   aws.String("aiff"),
	}

	_, err := b.bucket.PutObject(params)
	if err != nil {
		return errors.Wrap(err, "error occured uploading file to s3 bucket")
	}

	return nil
}
