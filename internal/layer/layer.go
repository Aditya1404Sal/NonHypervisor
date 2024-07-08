package layer

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"os/exec"
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

// RunCommand runs a shell command and returns any errors encountered
func RunCommand(command string) error {
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// CopyFiles copies files from the source to the destination
func CopyFiles(src, dest string) error {
	cmd := exec.Command("cp", "-r", src, dest)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
