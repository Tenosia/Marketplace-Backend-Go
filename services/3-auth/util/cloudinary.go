package util

import (
	"context"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/joho/godotenv"
	"os"
)

type Cloudinary struct {
	cld *cloudinary.Cloudinary
}

func NewCloudinary() *Cloudinary {
	godotenv.Load()
	cld, err := cloudinary.NewFromParams(
		os.Getenv("CLOUDINARY_CLOUD_NAME"),
		os.Getenv("CLOUDINARY_API_KEY"),
		os.Getenv("CLOUDINARY_API_SECRET"),
	)
	if err != nil {
		panic(err)
	}

	return &Cloudinary{cld: cld}
}

func (c *Cloudinary) UploadImage(ctx context.Context, file interface{}) (*uploader.UploadResult, error) {
	return c.cld.Upload.Upload(ctx, file, uploader.UploadParams{})
}
