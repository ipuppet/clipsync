package flags

import (
	"flag"

	"clipsync/internal/server"
)

var (
	Address    string
	ConfigPath string
)

const (
	portDefault = "8080"
	portUsage   = "The service listening port."

	addressDefault = "0.0.0.0"
	addressUsage   = "Set listen address."

	configPathDefault = "./config.json"
	configPathUsage   = "Set config file path."
)

func init() {
	flag.StringVar(&server.Port, "p", portDefault, portUsage)
	flag.StringVar(&Address, "addr", addressDefault, addressUsage)
	flag.StringVar(&ConfigPath, "config", configPathDefault, configPathUsage)

	flag.Parse()
}
