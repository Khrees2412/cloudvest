package utils

import (
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type UploadResponse struct {
	Url string
}

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
			Region:      aws.String(AWS_S3_REGION),
			Credentials: credentials.NewStaticCredentials(AWS_ACCESS_KEY, AWS_SECRET_KEY, ""),
		})
	if err != nil {
		panic(err)
	}
	return sess
}

// file, header, err := r.FormFile("file")
// if err != nil {
//     // Do your error handling here
//     return
// }
// defer file.Close()

// filename := header.Filename
func UploadFile(file *multipart.FileHeader) (*UploadResponse, interface{}) {
	tempfile, err := os.Open(file.Filename)
	if err != nil {
		return nil, err
	}
	defer tempfile.Close()

	// get the file size and read
	// the file content into a buffer
	var size = file.Size
	buffer := make([]byte, size)
	tempfile.Read(buffer)
	uploader := s3manager.NewUploader(sess)

	_, uploaderr := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(AWS_S3_BUCKET), // Bucket to be used
		Key:    aws.String(""),            // Name of the file to be saved
		Body:   tempfile,                  // File
	})
	if uploaderr != nil {
		return nil, err
	}
	return nil, nil
}
