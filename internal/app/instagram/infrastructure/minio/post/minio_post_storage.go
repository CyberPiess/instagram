package post

import (
	"log"

	"github.com/minio/minio-go"
)

const (
	bucketName = "instagram"
)

type MinioPostStorage struct {
	minioClient *minio.Client
}

func NewMinioPostStorage(minioClient *minio.Client) *MinioPostStorage {
	return &MinioPostStorage{minioClient: minioClient}
}

func (m *MinioPostStorage) UploadFile(image ImageDTO) {
	n, err := m.minioClient.FPutObject(bucketName, image.ObjectName, image.FilePath, minio.PutObjectOptions{ContentType: image.ContentType})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", image.ObjectName, n)
}
