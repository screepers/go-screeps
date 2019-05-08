package config

// Config Root .screeps.yml
type Config struct {
	Servers map[string]ServerConfig `yaml:"servers"`
	Configs map[string]AppConfig    `yaml:"configs"`
}

// AppConfig .screeps.yml app config
type AppConfig interface{}

// ServerConfig .screeps.yml server config
type ServerConfig struct {
	Host     string `yaml:"host"`
	Port     int16  `yaml:"port"`
	Secure   bool   `yaml:"secure"`
	Token    string `yaml:"token"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	PTR      bool   `yaml:"ptr"`
}
