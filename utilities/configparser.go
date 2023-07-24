package utilities

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Parameters struct {
	Port int
}

func ParseConfig() (p *Parameters) {
	path, err := os.Executable()
	if err != nil {
		log.Fatalf("Error getting executable path, %s", err)
	}
	executablePath := filepath.Dir(path)
	filePath := flag.String("config", executablePath+"\\config.yml", "Path of the configuration file in YAML format")

	if _, err := os.Stat(*filePath); os.IsNotExist(err) {
		file, err := os.Create(*filePath)
		if err != nil {
			log.Fatalf("Error creating config file, %s", err)
		}
		defer file.Close()

		_, err = file.WriteString("port: 8080")
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

	if err := viper.Unmarshal(&p); err != nil {
		log.Fatalf("Unable to decode into struct, %v\n", err)
	}

	return
}
