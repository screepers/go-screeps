package config

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v2"
)

// NewConfig load new Config
func NewConfig() *Config {
	paths := getDirs()
	for _, file := range paths {
		nc, err := loadConfig(file)
		if err != nil {
			continue
		}
		log.Printf("Loaded config from %s", file)
		return nc
	}
	return &Config{}
}

func loadConfig(file string) (*Config, error) {
	configFile, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	c := &Config{}
	err = yaml.Unmarshal(configFile, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func getDirs() []string {
	config1 := "config.yml"
	config2 := "config.yaml"
	screeps1 := ".screeps.yml"
	screeps2 := ".screeps.yaml"
	paths := make([]string, 0)
	if val := os.Getenv("SCREEPS_CONFIG"); val != "" {
		paths = append(paths, val)
	}
	ex, err := os.Executable()
	if err == nil {
		dir := filepath.Dir(ex)
		paths = append(paths, filepath.Join(dir, config1), filepath.Join(dir, config2))
	}
	paths = append(paths, screeps1)
	paths = append(paths, screeps2)

	if runtime.GOOS == "windows" {
		ad := os.Getenv("APPDATA")
		paths = append(paths, filepath.Join(ad, "screeps", "config.yml"))
		paths = append(paths, filepath.Join(ad, "screeps", "config.yaml"))
	} else {
		if val := os.Getenv("XDG_CONFIG_PATH"); val != "" {
			paths = append(paths, filepath.Join(val, "screeps", "config.yml"))
			paths = append(paths, filepath.Join(val, "screeps", "config.yaml"))
		}
		if val := os.Getenv("HOME"); val != "" {
			paths = append(paths, filepath.Join(val, ".config", "screeps", "config.yml"))
			paths = append(paths, filepath.Join(val, ".config", "screeps", "config.yaml"))
			paths = append(paths, filepath.Join(val, ".screeps.yml"))
			paths = append(paths, filepath.Join(val, ".screeps.yaml"))
		}
	}
	return paths
}
