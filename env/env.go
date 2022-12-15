package env

import (
	"fmt"
	"os"
)

type s3 struct {
	ACCESS_KEY      string
	SECRET_KEY      string
	REGION          string
	BUCKET          string
	BUCKET_ENDPOINT string
}

var S3Env *s3

func Load() {
	S3Env = new(s3)
	S3Env.initS3ENV()

	fmt.Println("ENV loaded!")
}

func (env *s3) initS3ENV() {
	env.ACCESS_KEY = os.Getenv("AWS_S3_ACCESS_KEY")
	env.SECRET_KEY = os.Getenv("AWS_S3_SECRET_KEY")
	env.REGION = os.Getenv("AWS_S3_REGION")
	env.BUCKET = os.Getenv("AWS_S3_BUCKET")
	env.BUCKET_ENDPOINT = os.Getenv("AWS_S3_BUCKET_ENDPOINT")
}
