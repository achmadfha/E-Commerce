package utils

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
	"io"
	"time"
)

func UploadToS3(file io.Reader) (string, error) {
	bucket := configData.DbConfig.AwsBucketS3
	key := configData.DbConfig.AwsAccessKey
	secret := configData.DbConfig.AwsSecretAccessKey

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-southeast-1"),
		Credentials: credentials.NewStaticCredentials(key, secret, ""),
	})
	if err != nil {
		return "", err
	}

	uploader := s3manager.NewUploader(sess)
	uuid := uuid.New()
	date := time.Now().Format("02-01-2006")
	fileName := fmt.Sprintf("%s-%s", uuid, date)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(fileName + ".jpeg"),
		Body:        file,
		Expires:     aws.Time(time.Now().Add(1000 * time.Second)),
		ContentType: aws.String("image/jpeg"),
	})
	if err != nil {
		return "", err
	}

	return "https://" + bucket + ".s3.ap-southeast-1.amazonaws.com/" + fileName + ".jpeg", nil
}
