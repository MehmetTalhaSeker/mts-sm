package config

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"

	"github.com/MehmetTalhaSeker/mts-sm/assets"
)

type Config struct {
	Rest struct {
		Host    string `yaml:"host"`
		Port    string `yaml:"port"`
		BaseURL string `yaml:"base_url"`
		Version string `yaml:"version"`
	} `yaml:"rest"`
	DB struct {
		Driver   string
		Host     string
		Port     string
		Name     string
		User     string
		Password string
		SSL      string
		TimeZone string
		Idle     int
		Open     int
	}
	Env      string `yaml:"env"`
	Security struct {
		Jwt struct {
			Exp string
			Key string
		}
	}
	Version bool `yaml:"version"`
	Minio   struct {
		Host                string
		Port                string
		Region              string
		Access              string
		Secret              string
		BucketName          string
		SupportedExtensions []string
	}
}

func Init() *Config {
	env := os.Getenv("ENV")
	if len(env) == 0 {
		env = "local"
	}

	path := fmt.Sprintf(`configs/env.%s.yaml`, env)

	file, err := assets.EmbeddedFiles.ReadFile(path)
	if err != nil {
		log.Fatal("fatal error config file: \n", err)
	}

	viper.AddConfigPath("./")
	viper.SetConfigName(fmt.Sprintf("env.%s", env))
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err = viper.ReadConfig(bytes.NewReader(file))
	if err != nil {
		log.Fatal("fatal error config file: \n", err)
	}

	var c Config

	err = viper.Unmarshal(&c)
	if err != nil {
		log.Fatal("fatal error config file: \n", err)
	}

	log.Println("Config initialize success!")

	return &c
}
