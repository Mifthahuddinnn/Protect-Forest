package config

import (
	"github.com/cloudinary/cloudinary-go"
	"log"
	"os"
)

func InitCloudinary() (*cloudinary.Cloudinary, error) {
	cloudinaryURL := os.Getenv("CLOUDINARY_URL")
	if cloudinaryURL == "" {
		log.Fatal("CLOUDINARY_URL is not set")
	}

	cloudinaryService, err := cloudinary.NewFromURL(cloudinaryURL)
	if err != nil {
		log.Printf("Failed to initialize Cloudinary: %v", err)
		return nil, err
	}

	return cloudinaryService, nil
}
