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

func Connect() (err error) {
	cld, err = cloudinary.NewFromURL(config.CloudinaryUrl)
	if err != nil {
		return
	}

	cld.Config.URL.Secure = true
	ctx = context.Background()
	return
}

func Upload(image *multipart.FileHeader, fileName string) (*StoredImage, error) {
	imageFile, _ := image.Open()
	defer imageFile.Close()

	uploadResult, err := cld.Upload.Upload(ctx, imageFile, uploader.UploadParams{
		PublicID: fileName,
		Folder:   "courses",
	})
	if err != nil {
		return nil, err
	}

	urlWithoutExtension := removeExtension(uploadResult.SecureURL)

	return &StoredImage{
		Url:          urlWithoutExtension,
		ThumbnailUrl: getThumbnailUrl(urlWithoutExtension),
	}, nil
}

func Rename(oldName, newName string) (*StoredImage, error) {
	updatedResult, err := cld.Upload.Rename(ctx, uploader.RenameParams{
		FromPublicID: oldName,
		ToPublicID:   newName,
	})
	if err != nil {
		return nil, err
	}

	return &StoredImage{
		Url:          updatedResult.SecureURL,
		ThumbnailUrl: getThumbnailUrl(updatedResult.SecureURL),
	}, nil
}

func Delete(fileName string) error {
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

func removeExtension(fileName string) string {
	return strings.Split(fileName, ".")[0]
}

func getThumbnailUrl(url string) string {
	return strings.Replace(url, "upload/", "upload/w_auto,h_300,c_fill,g_auto,f_auto/", 1)
}
