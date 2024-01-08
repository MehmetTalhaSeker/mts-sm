package fs

import (
	"context"
	"fmt"
	"github.com/MehmetTalhaSeker/mts-sm/internal/shared/config"
	"github.com/MehmetTalhaSeker/mts-sm/internal/shared/logg"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/zap"
	"io"
)

type MinioImpl struct {
	client     *minio.Client
	bucketName string
}

func Init(conf *config.Config) (*MinioImpl, error) {
	endpoint := fmt.Sprintf("%s:%s", conf.Minio.Host, conf.Minio.Port)

	FS, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(conf.Minio.Access, conf.Minio.Secret, ""),
		Secure: false,
	})

	if err != nil {
		logg.L.Error("err", zap.Error(err))
		return nil, err
	}

	logg.L.Info("File storage connection success!")

	return &MinioImpl{client: FS, bucketName: conf.Minio.BucketName}, err
}

func (m *MinioImpl) UploadImage(folderName, fileExt string, file io.Reader, size int64) (*string, error) {
	ctx := context.Background()

	createdName := prepareName(folderName, fileExt)
	object, err := m.client.PutObject(ctx, m.bucketName, createdName, file, size, minio.PutObjectOptions{})

	if err != nil {
		logg.L.Error("err", zap.Error(err))
		return nil, err
	}
	return &object.Key, nil
}

func prepareName(folderName, fileExt string) string {
	imageName := uuid.New().String() + fileExt
	prefix := fmt.Sprintf("%s/%s", folderName, imageName)
	return prefix
}
