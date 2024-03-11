package minio

import (
	"log"

	"github.com/minio/minio-go"
)

type MinioCred struct {
	Endpoint        string
	AccessKeyId     string
	SecretAccessKey string
	UseSSL          bool
}

func NewMinioConnection(cred MinioCred) (*minio.Client, error) {
	minioClient, err := minio.New(cred.Endpoint, cred.AccessKeyId, cred.SecretAccessKey, cred.UseSSL)
	if err != nil {
		log.Fatalln(err)
	}

	return minioClient, err
}
