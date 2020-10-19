package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type configer interface {
	unmarshalKey(key string, rawVal interface{}) error
	getString(key string) string
}

type config struct {
	internal *viper.Viper
}

func newConfig() (configer, error) {
	v := viper.New()
	v.SetConfigFile("config.toml")
	v.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("new config: reading in config: %w", err)
	}

	return &config{
		internal: v,
	}, nil
}

func (c *config) unmarshalKey(key string, rawVal interface{}) error {
	return c.internal.UnmarshalKey(key, rawVal)
}

func (c *config) getString(key string) string {
	return c.internal.GetString(key)
}
