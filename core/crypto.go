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
	defer Log.Sync()

	hash := sha256.New()

	if _, err := hash.Write(payload); err != nil {
		Log.Errorln("Cannot construct sha-256 hash")
	}

	return hash.Sum(nil)
}

func Sha512(payload []byte) []byte {
	defer Log.Sync()

	hash := sha512.New()

	if _, err := hash.Write(payload); err != nil {
		Log.Errorln("Cannot construct sha-512 hash")
	}

	return hash.Sum(nil)
}
