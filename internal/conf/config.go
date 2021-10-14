package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
)

type config struct {
	Site struct {
		Name string `toml:"name"`
		Port int    `toml:"port"`
	} `toml:"site"`
}

var conf *config

func Load() error {
	c := config{}

	_, err := toml.DecodeFile("./config/ddns.toml", &c)
	if err != nil {
		return errors.Wrap(err, "decode config file")
	}

	conf = &c
	return nil
}

// Get returns the config struct.
func Get() *config {
	return conf
}
