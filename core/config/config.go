package config

import (
	"os"

	"github.com/sunweiwe/paddle/pkg/config/db"
	"github.com/sunweiwe/paddle/pkg/config/redis"
	"github.com/sunweiwe/paddle/pkg/config/server"
	"github.com/sunweiwe/paddle/pkg/config/session"
	"gopkg.in/yaml.v3"
)

type Config struct {
	ServerConfig  server.Config  `yaml:"serverConfig"`
	DBConfig      db.Config      `yaml:"dbConfig"`
	KubeConfig    string         `yaml:"kubeconfig"`
	RedisConfig   redis.Redis    `yaml:"redisConfig"`
	SessionConfig session.Config `yaml:"sessionConfig"`
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
