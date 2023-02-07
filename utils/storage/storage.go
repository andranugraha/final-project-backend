package storage

import (
	"context"
	"errors"
	"final-project-backend/config"
	"mime/multipart"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

var (
	cld *cloudinary.Cloudinary
	ctx context.Context
)

type StoredImage struct {
	Url          string
	ThumbnailUrl string
}

type StorageUtil interface {
	Upload(*multipart.FileHeader, string) (*StoredImage, error)
	Rename(string, string) (*StoredImage, error)
	Delete(string) error
}

type storageUtilImpl struct{}

func NewStorageUtil() StorageUtil {
	return storageUtilImpl{}
}

func Connect() (err error) {
	cld, err = cloudinary.NewFromURL(config.CloudinaryUrl)
	if err != nil {
		return
	}

	cld.Config.URL.Secure = true
	ctx = context.Background()
	return
}

func (s storageUtilImpl) Upload(image *multipart.FileHeader, fileName string) (*StoredImage, error) {
	imageFile, _ := image.Open()
	defer imageFile.Close()

	uploadResult, err := cld.Upload.Upload(ctx, imageFile, uploader.UploadParams{
		PublicID: fileName,
		Folder:   "courses",
	})
	if err != nil {
		return nil, err
	}

	return &StoredImage{
		Url:          uploadResult.SecureURL,
		ThumbnailUrl: s.getThumbnailUrl(uploadResult.SecureURL),
	}, nil
}

func (s storageUtilImpl) Rename(oldName, newName string) (*StoredImage, error) {
	updatedResult, err := cld.Upload.Rename(ctx, uploader.RenameParams{
		FromPublicID: "courses/" + oldName,
		ToPublicID:   newName,
	})
	if err != nil {
		return nil, err
	}

	return &StoredImage{
		Url:          updatedResult.SecureURL,
		ThumbnailUrl: s.getThumbnailUrl(updatedResult.SecureURL),
	}, nil
}

func (s storageUtilImpl) Delete(fileName string) error {
	res, err := cld.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: "courses/" + fileName,
	})
	if err != nil {
		return err
	}

	if res.Result != "ok" {
		return errors.New("failed to delete image")
	}

	return nil
}

func (s *storageUtilImpl) getThumbnailUrl(url string) string {
	return strings.Replace(url, "upload/", "upload/w_auto,h_300,c_fill,g_auto,f_auto/", 1)
}
