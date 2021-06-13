package main

import (
	"bytes"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	AWSRegion     = "us-east-1"
	AWSBucketName = "sharrans-logs"
)

func newAWSSession() *session.Session {
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(AWSRegion),
		})
	if err != nil {
		panic(err)
	}
	return sess
}

func NewS3Uploader() *S3Uploader {
	return &S3Uploader{
		Session: newAWSSession(),
	}
}

type S3Uploader struct {
	Session *session.Session
}

// AddFileToS3 will upload a single file to S3, it will require a pre-built aws session
// and will set file info like content type and encryption on the uploaded file.
func (uploader *S3Uploader) AddFileToS3(fileDir string) error {

	// Open the file for use
	file, err := os.Open(fileDir)
	if err != nil {
		return err
	}
	defer file.Close()

	// Get file size and read the file content into a buffer
	fileInfo, _ := file.Stat()
	size := fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)

	// Config settings: this is where you choose the bucket, filename, content-type etc.
	// of the file you're uploading.
	_, err = s3.New(uploader.Session).PutObject(&s3.PutObjectInput{
		Bucket:        aws.String(AWSBucketName),
		Key:           aws.String(fileDir),
		Body:          bytes.NewReader(buffer),
		ContentLength: aws.Int64(size),
		ContentType:   aws.String(http.DetectContentType(buffer)),
	})
	return err
}
