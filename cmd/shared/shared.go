package shared

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

func Input(filePath string, config *config.Config) error {
	inputFile, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Unable to read file: %v", err)
		return err
	}
	err = yaml.Unmarshal(inputFile, config)
	if err != nil {
		fmt.Printf("Unable to unmarshal data: %v", err)
		return err
	}
	return nil
}
