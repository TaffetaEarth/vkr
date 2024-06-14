package minioclient

import (
	"context"

	"crud/app/utils"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const (
	endpoint = "minio:9000"
	accessKeyID = "nxKqBpa3FugO4XrVUIYv"
	secretAccessKey = "FgiHh62MPaNNSMLPl4yovHVp4jC8uw1O8gWX5tvN"
	bucketName = "music"
)

func SetupMinioClient() *minio.Client {
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})
	utils.CheckErr(err)
	bucketExists, err := minioClient.BucketExists(context.Background(), bucketName);
	utils.CheckErr(err)
	if !bucketExists {
		err := minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
		utils.CheckErr(err)
	}

	return minioClient
}
