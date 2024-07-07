package config

import (
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

// Config represents the structure of the build.yaml file
type Config struct {
	BaseImage string  `yaml:"base_image"`
	Layers    []Layer `yaml:"layers"`
}

// Layer represents a single layer in the configuration
type Layer struct {
	Run  string `yaml:"run,omitempty"`
	Copy struct {
		Src  string `yaml:"src,omitempty"`
		Dest string `yaml:"dest,omitempty"`
	} `yaml:"copy,omitempty"`
}

// ParseConfig parses the YAML configuration file
func ParseConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
