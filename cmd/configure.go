/*
Copyright © 2023 TechSquidTV
*/
package cmd

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"github.com/techsquidtv/uhs-cli/cmd/common"
	"github.com/techsquidtv/uhs-cli/models"
	"github.com/techsquidtv/uhs-cli/models/services"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure your UHS instance",
	Long:  `Customize and configure your desired services for your UHS instance.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Create a new instance of the UHSConfig struct
		uhsConfig := models.UHSConfig{
			Services: make(services.ServicesConfig),
		}
		// Return each key from the DefaultServiceConfig() config map
		serviceNames := make([]string, 0, len(models.DefaultServiceConfig()))
		for k := range models.DefaultServiceConfig() {
			serviceNames = append(serviceNames, k)
		}
		// Prompt user to select services to enable
		serviceSelectPrompt := &survey.MultiSelect{
			Message: "Select services to enable:",
			Options: serviceNames,
		}
		var selectedServices []string
		err := survey.AskOne(serviceSelectPrompt, &selectedServices)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		// Validate selected services
		serviceListString := ""
		for _, service := range selectedServices {
			serviceListString += fmt.Sprintf("  - %v\n", service)
		}

		validateSelectionPrompt := &survey.Confirm{
			Message: fmt.Sprintf("You have selected the following services:\n%v\n Is this correct?", serviceListString),
		}
		var validateSelection bool
		survey.AskOne(validateSelectionPrompt, &validateSelection)
		// If selection is not valid, exit
		if !validateSelection {
			fmt.Println("Exiting...")
			return
		}
		// Execute configuration for each selected service
		for _, serviceName := range selectedServices {
			fmt.Println("Configuring " + serviceName + "...")
			service := models.DefaultServiceConfig()[serviceName]
			config := service.Configure()
			uhsConfig.Services[serviceName] = config
		}
		// Output
		outputFile, err := cmd.Flags().GetString("output")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		err = common.Output(outputFile, &uhsConfig)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
	configureCmd.PersistentFlags().StringP("output", "o", "", "Output file path")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configureCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configureCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
