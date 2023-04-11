//go:build windows

package conf

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Params struct {
	Port int
}

func NewParams() (p *Params) {
	ex, err := os.Executable()
	if err != nil {
		log.Fatalf("Error getting executable path, %s", err)
	}
	executablePath := filepath.Dir(ex)

	filePath := flag.String("config", executablePath+"\\config.yml", "Path of the configuration file in YAML format")

	if _, err := os.Stat(*filePath); os.IsNotExist(err) {
		f, err := os.Create(*filePath)
		if err != nil {
			log.Fatalf("Error creating config file, %s", err)
		}
		defer f.Close()

		_, err = f.WriteString("port: 8080")
		if err != nil {
			log.Fatalf("Error writing config file, %s", err)
		}

		log.Println("Default config file created in: " + *filePath)
	}

	flag.Parse()
	viper.SetConfigFile(*filePath)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err = viper.Unmarshal(&p)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v\n", err)
	}

	return
}
