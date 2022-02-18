package utils

import (
	"fmt"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var (
	AWS_S3_REGION  = os.Getenv("AWS_REGION")
	AWS_S3_BUCKET  = os.Getenv("AWS_BUCKET")
	AWS_ACCESS_KEY = os.Getenv("AWS_ACCESS_KEY")
	AWS_SECRET_KEY = os.Getenv("AWS_SECRET_KEY")
)

var sess = connectAWS()

func connectAWS() *session.Session {
	sess, err := session.NewSession(
		&aws.Config{
			Region:      aws.String("eu-west-2"),
			Credentials: credentials.NewStaticCredentials(AWS_ACCESS_KEY, AWS_SECRET_KEY, ""),
		})
	if err != nil {
		panic(err)
	}
	return sess
}

func UploadFile(file *multipart.FileHeader) (*s3manager.UploadOutput, interface{}) {
	buffer, err := file.Open()

	if err != nil {
		fmt.Println(err)
	}
	defer buffer.Close()

	uploader := s3manager.NewUploader(sess)

	data, uploaderr := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("risevest"),    // Bucket to be used
		Key:    aws.String(file.Filename), // Name of the file to be saved
		Body:   buffer,                    // File
	})
	if uploaderr != nil {
		return nil, uploaderr
	}
	return data, nil
}
