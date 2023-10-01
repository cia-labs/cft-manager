package appconf

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type AppConf struct {
	APIBaseURL    string `yaml:"api_base_url"`
	APITunnelList string `yaml:"api_tunnel_list"`
	APIKey        string `yaml:"api_key"`
	AccountID     string `yaml:"account_id"`
}

var AppConfig AppConf

func LoadAppConf(configFile string) {
	// read the YAML file
	yamlFile, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalln("Failed to read YAML file: ", err)
	}

	// unmarshal the YAML data into an AppConf struct
	err = yaml.Unmarshal(yamlFile, &AppConfig)
	if err != nil {
		log.Fatalln("Failed to unmarshal YAML data: ", err)
	}

}
