package image

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"mime/multipart"
)

func CheckImg(f *multipart.FileHeader) bool {
	img, _ := f.Open()
	_, _, err := image.Decode(img)
	if err != nil {
		_ = img.Close()
		return false
	}
	_ = img.Close()
	return true
}
