package utils

import (
	"encoding/base64"
	"github.com/skip2/go-qrcode"
)

func GenerateQRCode(content string) (string, error) {
	png, err := qrcode.Encode(content, qrcode.Medium, 256)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(png), nil
}
