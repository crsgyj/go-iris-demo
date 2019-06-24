package config

import (
	"fmt"

	"github.com/pelletier/go-toml"
)

var (
	// Conf - config object
	Conf = newConf()
)

// Config - config struct
type Config struct {
	App struct {
		Port           string
		AllowedOrigins []string
	}
	Redis struct {
		Addr     string
		Password string
	}
	BaiduFY struct {
		Appid string
		Key   string
	}
}

func newConf() *Config {
	conf, err := toml.LoadFile("D:\\crsgy\\project\\go_project\\src\\comm-filter\\config\\development.toml")
	if err != nil {
		fmt.Println("TomlError ", err.Error())
	}

	v := &Config{}
	conf.Unmarshal(v)
	return v
}
