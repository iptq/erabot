package erabot

import "github.com/BurntSushi/toml"

type Config struct {
	Token   string
	ModConf map[string]interface{}
}

var DefaultConfig = Config{
	Token: "<your discord bot token>",
}

func LoadConfig(contents string) (config Config, err error) {
	_, err = toml.Decode(contents, &config)
	return
}

func (config *Config) AddModConfig(name string, conf interface{}) {
	config.ModConf[name] = conf
}
