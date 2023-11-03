package helpers

import (
	"context"
	"fmt"
	"mime/multipart"
	"perpustakaan/config"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func (h *helper) UploadImage(folder string, file multipart.File) (string) {
	fmt.Println("In Cloud Upload Image, formFile type : ", &file)
	cfg := config.LoadServerConfig()
	cld, err := cloudinary.NewFromURL(cfg.CLOUD_URL)

	if err != nil || file == nil {
		return ""
	}

	ctx := context.Background()

	upload, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
		Folder: folder,
	})

	return upload.SecureURL
}