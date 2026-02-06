package utils

import (
	"bytes"
	"encoding/base64"
	"strings"
	"context"
	"errors"
	"regexp"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func DecodeBase64Image(b64 string) ([]byte, error) {
	if strings.Contains(b64, ",") {
		parts := strings.SplitN(b64, ",", 2)
		b64 = parts[1]
	}

	return base64.StdEncoding.DecodeString(b64)
}

func UploadBase64Image(
	ctx context.Context,
	client *s3.Client,
	bucket string,
	key string,
	b64 string,
	contentType string,
) error {

	data, err := DecodeBase64Image(b64)
	if err != nil {
		return err
	}

	_, err = client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      &bucket,
		Key:         &key,
		Body:        bytes.NewReader(data),
		ContentType: &contentType,
	})

	return err
}

func ExtractUUIDFromImageURL(url string) (string, error) {
	re := regexp.MustCompile(`([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}_image)`)
	matches := re.FindStringSubmatch(url)

	if len(matches) < 2 {
		return "", errors.New("uuid not found in image URL")
	}

	return matches[1], nil
}
