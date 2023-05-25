package common

import (
	"fmt"
	"os"

	"github.com/techsquidtv/uhs-cli/models/config"
	"gopkg.in/yaml.v3"
)

func Output(filePath string, config *config.Config) error {
	yamlConfig, err := yaml.Marshal(config)
	if err != nil {
		return err
	}
	if filePath == "" {
		fmt.Println(string(yamlConfig))
		return nil
	} else {
		err = os.WriteFile(filePath, yamlConfig, 0644)
		if err != nil {
			return err
		}
		fmt.Println("Configuration complete! Your configuration file has been saved to " + filePath + ".")
	}
	return nil
}
