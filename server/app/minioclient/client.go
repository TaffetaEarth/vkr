package minioclient

import (
	"log"
	"context"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const (
	endpoint = "minio:9000"
	accessKeyID = "9C2QBjaBZ2eQAzKy0DQr"
	secretAccessKey = "1voNvvXh1axgMRzR5ykNwd8YrDXY4uuQMdYmS8Q4"
	bucketName = "music"
)

func SetupMinioClient() *minio.Client {
	
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})

	checkErr(err)

	bucketExists, err := minioClient.BucketExists(context.Background(), bucketName);
	
	checkErr(err)

	if !bucketExists {
		err := minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
		checkErr(err)
	}

	return minioClient
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
} 