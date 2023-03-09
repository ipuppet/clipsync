package flags

import (
	"flag"

	"clipsync/internal/server"
)

var Address string

const (
	portDefault = "8080"
	portUsage   = "The service listening port."

	addressPathDefault = "0.0.0.0"
	addressPathUsage   = "Set listen address."
)

func init() {
	flag.StringVar(&server.Port, "p", portDefault, portUsage)

	flag.StringVar(&Address, "addr", addressPathDefault, addressPathUsage)

	flag.Parse()
}
