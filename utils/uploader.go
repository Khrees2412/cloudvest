package utils

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
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
			Region:      aws.String(AWS_S3_REGION),
			Credentials: credentials.NewStaticCredentials(AWS_ACCESS_KEY, AWS_SECRET_KEY, ""),
		})
	if err != nil {
		panic(err)
	}
	return sess
}

// func uploadFileToS3(s *session.Session, fileName string) error {

//     // open the file for use
//     file, err := os.Open(fileName)
//     if err != nil {
//         return err
//     }
//     defer file.Close()

//     // get the file size and read
//     // the file content into a buffer
//     fileInfo, _ := file.Stat()
//     var size = fileInfo.Size()
//     buffer := make([]byte, size)
//     file.Read(buffer)

//     // config settings: this is where you choose the bucket,
//     // filename, content-type and storage class of the file
//     // you're uploading
//     e, s3err := s3.New(s).PutObject(&s3.PutObjectInput{
//         Bucket:               aws.String(S3_BUCKET),
//         Key:                  aws.String(fileName),
//         ACL:                  aws.String("private"),
//         Body:                 bytes.NewReader(buffer),
//         ContentLength:        aws.Int64(size),
//         ContentType:          aws.String(http.DetectContentType(buffer)),
//         ContentDisposition:   aws.String("attachment"),
//         ServerSideEncryption: aws.String("AES256"),
//         StorageClass:         aws.String("INTELLIGENT_TIERING"),
//     })

//     return s3err
// }

// file, header, err := r.FormFile("file")
// if err != nil {
//     // Do your error handling here
//     return
// }
// defer file.Close()

// filename := header.Filename

// uploader := s3manager.NewUploader(sess)

// _, err = uploader.Upload(&s3manager.UploadInput{
//     Bucket: aws.String(AWS_BUCKET), // Bucket to be used
//     Key:    aws.String(""),      // Name of the file to be saved
//     Body:   file,                      // File
// })
