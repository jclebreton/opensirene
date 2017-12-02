package files

import (
	"crypto/sha256"
	"fmt"
	"hash"
	"io"
	"os"

	"github.com/pkg/errors"
)

// CalculateChecksum generates the checksum of the file
func CalculateChecksum(path string, checksumType string) (string, error) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return "", err
	}

	var hasher hash.Hash
	switch checksumType {
	case "sha256":
		hasher = sha256.New()
	default:
		return "", errors.New("unknown checksum type")
	}
	_, err = io.Copy(hasher, f)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hasher.Sum(nil)), nil
}
