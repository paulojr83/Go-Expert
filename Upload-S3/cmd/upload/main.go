package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
	"os"
	"sync"
)

var (
	s3Client *s3.S3
	s3Bucket string
	wg       sync.WaitGroup
)

func init() {
	sess, err := session.NewSession(
		&aws.Config{
			Credentials: credentials.NewStaticCredentials(
				"**********", "**************", ""),
			Region: aws.String("us-east-1"),
		},
	)

	if err != nil {
		panic(err)
	}

	s3Client = s3.New(sess)
	s3Bucket = "test-upload-go"
}

func main() {
	path := "./tmp"
	dir, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer dir.Close()
	uploadControl := make(chan struct{}, 100)
	errorFileUploadControl := make(chan string, 10)

	go func() {
		for {
			select {
			case filename := <-errorFileUploadControl:
				uploadControl <- struct{}{}
				wg.Add(1)
				go uploadFile(path, filename, uploadControl, errorFileUploadControl)
			}
		}
	}()
	for {
		files, err := dir.Readdir(1)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("Error reading directory: %s\n", err)
			continue
		}

		wg.Add(1)
		uploadControl <- struct{}{}
		go uploadFile(path, files[0].Name(), uploadControl, errorFileUploadControl)
	}
	wg.Wait()
}

func uploadFile(path, filename string, uploadControl <-chan struct{}, errorFileUploadControl chan<- string) {
	defer wg.Done()
	completeFilename := fmt.Sprintf("%s/%s", path, filename)
	fmt.Printf("Uploading file %s to bucket %s\n", filename, s3Bucket)

	file, err := os.Open(completeFilename)
	if err != nil {
		fmt.Printf("Error opening file %s\n", completeFilename)
		<-uploadControl //esvazia o canal
		errorFileUploadControl <- filename
		return
	}
	defer file.Close()
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Body:   file,
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(filename),
	})
	if err != nil {
		fmt.Printf("Error upload file %s\nMessage error: %s", filename, err)
		<-uploadControl //esvazia o canal
		errorFileUploadControl <- filename
		return
	}

	fmt.Printf("File %s uploaded successfully\n", filename)
	<-uploadControl //esvazia o canal
}
