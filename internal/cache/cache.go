package cache

import (
	"os"
	"path/filepath"
)

const cacheDir = "/tmp/mockdocker_cache"

// CheckCache checks if a layer is present in the cache
func CheckCache(layerHash string) bool {
	_, err := os.Stat(filepath.Join(cacheDir, layerHash))
	return !os.IsNotExist(err)
}

// StoreLayer stores a layer in the cache
func StoreLayer(layerHash string, layerData []byte) error {
	cachePath := filepath.Join(cacheDir, layerHash)
	return os.WriteFile(cachePath, layerData, 0644)
}
