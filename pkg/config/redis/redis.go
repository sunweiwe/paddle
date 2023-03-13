package redis

type Redis struct {
	Protocol string `yaml:"protocol"`
	Address  string `yaml:"address"`
	Password string `yamlL:"password"`
	DB       uint8  `yaml:"db"`
}
