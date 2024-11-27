package server

import (
	"clipsync/internal/flags"
	"clipsync/internal/pkg"
)

type Aes struct {
	Key string `json:"key"`
	IV  string `json:"iv"`
}

type Config struct {
	Aes Aes `json:"aes"`
}

var config *Config

func init() {
	pkg.GetConfig(flags.ConfigPath, &config)
}
