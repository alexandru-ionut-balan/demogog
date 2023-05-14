package core

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
)

func Base64(payload []byte) string {
	return base64.StdEncoding.EncodeToString(payload)
}

func Sha256(payload []byte) []byte {
	hash := sha256.New()

	hash.Write(payload)

	return hash.Sum(nil)
}

func Sha512(payload []byte) []byte {
	hash := sha512.New()

	hash.Write(payload)

	return hash.Sum(nil)
}
