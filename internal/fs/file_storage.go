package fs

import (
	"context"
	"fmt"
	"github.com/MehmetTalhaSeker/mts-sm/internal/shared/config"
	"github.com/MehmetTalhaSeker/mts-sm/internal/shared/logg"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
)

type IFileStorage interface {
	UploadImage(folderName, fileExt string, file io.Reader, size int64) (*string, error)
}

type minioFS struct {
	client     *minio.Client
	bucketName string
}

func New(conf *config.Config) (IFileStorage, error) {
	endpoint := fmt.Sprintf("%s:%s", conf.Minio.Host, conf.Minio.Port)

	FS, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(conf.Minio.Access, conf.Minio.Secret, ""),
		Secure: false,
	})

	ctx := context.Background()
	exists, err := FS.BucketExists(ctx, conf.Minio.BucketName)
	if err != nil {
		return nil, err
	}

	if !exists {
		err = FS.MakeBucket(ctx, conf.Minio.BucketName, minio.MakeBucketOptions{})
		if err != nil {
			return nil, err
		}
	}

	logg.L.Info("File storage connection success!")

	return &minioFS{client: FS, bucketName: conf.Minio.BucketName}, err
}

func (m *minioFS) UploadImage(folderName, fileExt string, file io.Reader, size int64) (*string, error) {
	ctx := context.Background()

	createdName := prepareName(folderName, fileExt)
	object, err := m.client.PutObject(ctx, m.bucketName, createdName, file, size, minio.PutObjectOptions{})

	if err != nil {
		return nil, err
	}
	return &object.Key, nil
}

func prepareName(folderName, fileExt string) string {
	imageName := uuid.New().String() + fileExt
	prefix := fmt.Sprintf("%s/%s", folderName, imageName)
	return prefix
}
