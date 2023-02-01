package minio

import (
	"context"
	"github.com/hcdoit/tiktok/pkg/consts"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var MINIO *minio.Client

func Init() {
	client, err := minio.New(consts.MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(consts.MinioAccessKeyId, consts.MinioSecretAccessKey, ""),
		Secure: consts.MinioUseSSL,
	})
	if err != nil {
		panic(err)
	}

	MINIO = client

	err = InitBucket(consts.MinioVideoBucketName)
	if err != nil {
		panic(err)
	}

}
func InitBucket(bucketName string) error {

	location := "beijing"
	ctx := context.Background()

	err := MINIO.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		exists, errBucketExists := MINIO.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			return nil
		} else {
			return err
		}
	}

	err = MINIO.SetBucketPolicy(ctx, bucketName, consts.MinioPolicy)
	if err != nil {
		return err
	}
	return nil
}
