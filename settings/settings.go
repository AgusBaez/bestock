package settings

import (
	_ "embed"

	"gopkg.in/yaml.v3"
)

// Go Embed a la hora de iniciar la aplicacion todo lo que se encuentre en el archivo yalm lo envia hacia la variable "settingsFile"

//go:embed settings.yaml
var settingsFile []byte

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type Settings struct {
	Port string         `yaml:"port"`
	DB   DatabaseConfig `yaml:"database"`
}

func New() (*Settings, error) {
	var s Settings

	//PACKETE: yaml.Unmarshal descodifica el primer documento encontrado en el byte de entrada y asigna los valores descodificados al valor de salida.
	err := yaml.Unmarshal(settingsFile, &s)

	// nunca puede ser nil el parametro de salida
	if err != nil { 
		return nil, err
	}

	return &s, nil
}
