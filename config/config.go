package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	Port string `mapstructure:"SERVER_PORT"`
	URL  string `mapstructure:"URL"`
}

type Config struct {
	ServerConfig `mapstructure:",squash"`
}

var Global Config

func Load(env, configFolder string) error {
	v := viper.New()
	v.SetConfigName(".default")
	v.SetConfigType("env")

	if err := v.ReadInConfig(); err != nil {
		return err
	}

	if env != "" {
		env = strings.ToLower(env)
		v.SetConfigName(fmt.Sprintf(".%s", env))
		v.MergeInConfig()
	}
	v.SetConfigName(".local")
	v.MergeInConfig()

	v.AutomaticEnv()

	if err := v.Unmarshal(&Global); err != nil {
		return err
	}

	return nil

}
