package filesystem

import (
	"bytes"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/jonnypillar/somniloquy/configs"
	"github.com/pkg/errors"
)

// S3Bucket ...
type S3Bucket struct {
	bucket     *s3.S3
	bucketName string
	filePath   string
}

// NewS3Bucket ...
func NewS3Bucket(config *config.ServiceConfig) (*S3Bucket, error) {
	fmt.Println("Initilising S3 Saver")

	creds := credentials.NewCredentials(&credentials.SharedCredentialsProvider{})
	_, err := creds.Get()
	if err != nil {
		return nil, errors.Wrap(err, "error occured creating AWS credentials")
	}
	cfg := aws.NewConfig().WithRegion(config.AWSRegion).WithCredentials(creds)
	svc := s3.New(session.New(), cfg)

	return &S3Bucket{
		bucket:     svc,
		bucketName: config.AWSUploadS3BucketName,
		filePath:   config.FLACRecordingFilePath,
	}, nil
}

// Save ...
func (b S3Bucket) Save(filename string, by *bytes.Buffer) error {
	data := by.Bytes()
	fileBytes := bytes.NewReader(data)
	filePath := b.filePath + filename

	params := &s3.PutObjectInput{
		Bucket:        aws.String(b.bucketName),
		Key:           aws.String(filePath),
		Body:          fileBytes,
		ContentLength: aws.Int64(int64(len(data))),
		ContentType:   aws.String("aiff"),
	}

	_, err := b.bucket.PutObject(params)
	if err != nil {
		return errors.Wrap(err, "error occured uploading file to s3 bucket")
	}

	fmt.Println("Saved file to S3:", filename)

	return nil
}

// Read ...
func (S3Bucket) Read() ([]os.FileInfo, error) {
	return nil, nil
}
