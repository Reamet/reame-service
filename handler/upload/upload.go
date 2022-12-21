package upload

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"log"
	"reame-service/env"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var SUPPORT_MIME_TYPE = []string{"image/png", "image/jpg", "image/jpeg"}

func AWSConnect() (*session.Session, error) {
	fmt.Println(env.S3Env.ACCESS_KEY)
	creds := credentials.NewStaticCredentials(env.S3Env.ACCESS_KEY, env.S3Env.SECRET_KEY, "")
	_, err := creds.Get()
	if err != nil {
		return nil, err
	}
	sess, err := session.NewSession(
		&aws.Config{
			Region:      aws.String(env.S3Env.REGION),
			Endpoint:    aws.String(env.S3Env.BUCKET_ENDPOINT),
			Credentials: creds,
		},
	)
	if err != nil {
		return nil, err
	}
	return sess, nil
}

func AWSUpload(base64File string, filename string) (*s3manager.UploadOutput, error) {

	var mimeType = ""

	for _, mime := range SUPPORT_MIME_TYPE {
		if strings.Contains(base64File, mime) {
			mimeType = mime
			break
		}
	}

	if mimeType == "" {
		err := fmt.Errorf("mime: not support")
		log.Println(err)
		return nil, err
	}

	b64data := base64File[strings.IndexByte(base64File, ',')+1:]

	decode, err := base64.StdEncoding.DecodeString(b64data)
	if err != nil {
		return nil, err
	}

	sess, err := AWSConnect()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	uploader := s3manager.NewUploader(sess)

	uploadOutput, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(env.S3Env.BUCKET),
		Key:         aws.String(filename),
		Body:        bytes.NewReader(decode),
		ContentType: aws.String(mimeType),
		ACL:         aws.String("public-read"),
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return uploadOutput, err
}
