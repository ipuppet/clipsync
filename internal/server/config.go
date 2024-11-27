package server

import "clipsync/internal/pkg"

type Aes struct {
	Key string `json:"key"`
	IV  string `json:"iv"`
}

type Config struct {
	Aes Aes `json:"aes"`
}

var config *Config

func init() {
	pkg.GetConfig("config/config.json", &config)
}
