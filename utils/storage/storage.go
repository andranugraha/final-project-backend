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

type StoredImage struct {
	Url          string
	ThumbnailUrl string
}

func driver() (*cloudinary.Cloudinary, context.Context) {
	cld, _ := cloudinary.NewFromURL(config.CloudinaryUrl)
	cld.Config.URL.Secure = true
	ctx := context.Background()
	return cld, ctx
}

func Upload(image *multipart.FileHeader, fileName string) (*StoredImage, error) {
	imageFile, _ := image.Open()
	defer imageFile.Close()

	cld, ctx := driver()

	uploadResult, err := cld.Upload.Upload(ctx, imageFile, uploader.UploadParams{
		PublicID: fileName,
		Folder:   "courses",
	})
	if err != nil {
		return nil, err
	}

	return &StoredImage{
		Url:          uploadResult.SecureURL,
		ThumbnailUrl: getThumbnailUrl(uploadResult.SecureURL),
	}, nil
}

func Rename(oldName, newName string) (*StoredImage, error) {
	cld, ctx := driver()

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
	cld, ctx := driver()

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

func getThumbnailUrl(url string) string {
	return strings.Replace(url, "upload/", "upload/w_300,h_300,c_fill,g_auto/", 1)
}
