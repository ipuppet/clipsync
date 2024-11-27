package pkg

import (
	"encoding/json"
	"log"
	"os"
)

func GetConfig(path string, v interface{}) {
	configString, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := json.Unmarshal(configString, &v); err != nil {
		log.Fatal(err.Error())
	}
}
