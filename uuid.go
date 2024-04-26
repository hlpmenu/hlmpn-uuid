package gopkg.hlmpn.dev/uuid

import (
	"crypto/rand"
	"encoding/hex"
)

type UUID string

var err error

func New() (string, error) {
	b := make([]byte, 16)
	_, err = rand.Read(b)
	if err != nil {
		return "", err
	}
	uuid := hex.EncodeToString(b)
	// Insert hyphens after every 8 characters
	uuidWithHyphens := uuid[0:8] + "-" + uuid[8:16] + "-" + uuid[16:24] + "-" + uuid[24:]

	return string(uuidWithHyphens), nil
}
