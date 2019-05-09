package config

// Config Root .screeps.yml
type Config struct {
	Servers map[string]*ServerConfig `yaml:"servers"`
	Configs map[string]interface{}   `yaml:"configs"`
}

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
