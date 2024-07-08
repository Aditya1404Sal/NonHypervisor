package builder

import (
	"Nonhypervisor/internal/cache"
	"Nonhypervisor/internal/config"
	"Nonhypervisor/internal/layer"
	"fmt"
)

// BuildImage is supposed to build an image based on the configuration
func BuildImage(cfg *config.Config) error {
	fmt.Printf("Building image from base: %s\n", cfg.BaseImage)
	// BaseImage scaffolding is still incomplete.
	// OCI Image compliance is still incomplete.
	for _, l := range cfg.Layers {
		layerHash, err := layer.HashLayer(l)
		if err != nil {
			return fmt.Errorf("error hashing layer: %v", err)
		}

		if cache.CheckCache(layerHash) {
			fmt.Printf("Layer %s found in cache, skipping...\n", layerHash)
			continue
		}

		if l.Run != "" {
			fmt.Printf("Running command: %s\n", l.Run)
			err = layer.RunCommand(l.Run)
			if err != nil {
				return fmt.Errorf("error running command: %v", err)
			}
		} else if l.Copy.Src != "" && l.Copy.Dest != "" {
			fmt.Printf("Copying files from %s to %s\n", l.Copy.Src, l.Copy.Dest)
			err = layer.CopyFiles(l.Copy.Src, l.Copy.Dest)
			if err != nil {
				return fmt.Errorf("error copying files: %v", err)
			}
		}

		// Simulate storing the layer in cache
		err = cache.StoreLayer(layerHash, []byte("layer data"))
		if err != nil {
			return fmt.Errorf("error storing layer in cache: %v", err)
		}
	}

	return nil
}
