//go:build windows

package conf

import (
	"flag"
	"github.com/spf13/viper"
	"log"
	"os"
)

type Params struct {
	Port int
}

func NewParams() (p *Params) {
	filePath := flag.String("config", os.Getenv("PROGRAMFILES") + "\\wmi-rest\\config.yml", "Path of the configuration file in YAML format")
	flag.Parse()

	if _, err := os.Stat(*filePath); os.IsNotExist(err) {
		log.Fatalf("Configuration file: %s does not exist, %v\n", *filePath, err)
	}

	viper.SetConfigFile(*filePath)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&p)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v\n", err)
	}

	return
}
