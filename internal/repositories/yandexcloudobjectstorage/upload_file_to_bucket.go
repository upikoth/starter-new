package yandexcloudobjectstorage

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pkg/errors"
	"github.com/upikoth/starter-new/internal/model"
)

func (s *YandexCloudObjectStorage) UploadFileToBucket(
	ctx context.Context,
	req model.UploadFileToBucketRequest,
) error {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:   aws.String("ru-central1"),
			Endpoint: aws.String("https://storage.yandexcloud.net"),
			Credentials: credentials.NewStaticCredentialsFromCreds(credentials.Value{
				AccessKeyID:     req.AccessKey,
				SecretAccessKey: req.SecretKey,
			}),
		},
	}))

	storage := s3.New(sess)

	_, err := storage.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket: aws.String(req.BucketName),
		Key:    aws.String(req.FileName),
		Body:   bytes.NewReader(req.FileContent),
	})

	return errors.WithStack(err)
}
