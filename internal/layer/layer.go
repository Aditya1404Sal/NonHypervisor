package layer

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
)

// HashLayer generates a SHA-256 hash for a layer
func HashLayer(layer interface{}) (string, error) {
	hash := sha256.New()
	_, err := io.WriteString(hash, fmt.Sprintf("%v", layer))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}
