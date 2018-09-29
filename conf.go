package erabot

import "github.com/BurntSushi/toml"

type Config struct {
	modcfg map[string]interface{}
}

var DefaultConfig = Config{}

func LoadConfig(contents string) (config Config, err error) {
	_, err = toml.Decode(contents, &config)
	return
}
