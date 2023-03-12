package config

import (
	"os"

	"github.com/sunweiwe/paddle/pkg/config/db"
	"github.com/sunweiwe/paddle/pkg/config/server"
	"gopkg.in/yaml.v3"
)

type Config struct {
	ServerConfig server.Config `yaml:"serverConfig"`
	DBConfig     db.Config     `yaml:"dbConfig"`
	KubeConfig   string        `yaml:"kubeconfig"`
}

func LoadConfig(path string) (*Config, error) {
	var config Config
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
