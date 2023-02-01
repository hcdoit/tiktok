package minio

import (
	"bytes"
	"context"
	"github.com/hcdoit/tiktok/pkg/consts"
	"github.com/minio/minio-go/v7"
)

func SaveVideo(ctx context.Context, fileName string, fileData []byte) (localURL string, url string, err error) {
	reader := bytes.NewReader(fileData)
	size := int64(len(fileData))
	_, err = MINIO.PutObject(ctx, consts.MinioVideoBucketName, fileName, reader, size, minio.PutObjectOptions{
		ContentType: "video/mp4",
	})
	if err != nil {
		return "", "", err
	}
	localURL = "http://" + consts.MinioLocalAccessURL + "/" + consts.MinioVideoBucketName + "/" + fileName
	url = "http://" + consts.MinioAccessURL + "/" + consts.MinioVideoBucketName + "/" + fileName
	return localURL, url, nil
}

func SaveCover(ctx context.Context, fileName string, fileData []byte) (string, error) {
	reader := bytes.NewReader(fileData)
	size := int64(len(fileData))
	_, err := MINIO.PutObject(ctx, consts.MinioVideoBucketName, fileName, reader, size, minio.PutObjectOptions{
		ContentType: "image/jpeg",
	})
	if err != nil {
		return "", err
	}
	fileUrl := "http://" + consts.MinioAccessURL + "/" + consts.MinioVideoBucketName + "/" + fileName
	return fileUrl, nil
}
