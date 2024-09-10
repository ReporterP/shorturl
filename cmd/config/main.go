package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)
type AppConfig struct {
    AddressAndPort string `yaml:"addressandport"`
    BaseAddress string `yaml:"baseaddress"`
}
func (config *AppConfig) ReadYaml(filepath string) {
	yamlFile, err := os.ReadFile(filepath)
    if err != nil {
		fmt.Printf("Error Reading Config file with path: %v\n", filepath)
    }
    yaml.Unmarshal(yamlFile, config)

}

