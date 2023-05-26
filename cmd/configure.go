/*
Copyright © 2023 TechSquidTV
*/
package cmd

import (
	"fmt"
	"os"
	"sort"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"github.com/techsquidtv/uhs-cli/cmd/common"
	configCommon "github.com/techsquidtv/uhs-cli/models/common"
	"github.com/techsquidtv/uhs-cli/models/config"
	"github.com/techsquidtv/uhs-cli/models/services"
)

// Return each key from the DefaultServiceConfig() config map
var serviceNames = make([]string, 0, len(config.DefaultServiceConfig()))

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:       "configure",
	Short:     "Configure your UHS instance",
	Long:      `Customize and configure your desired services for your UHS instance.`,
	ValidArgs: serviceNames,
	Run: func(cmd *cobra.Command, args []string) {
		// Get selected services from args
		var selectedServices []string
		if len(args) > 0 {
			selectedServices = args
		}
		// Set services to default
		for k := range config.DefaultServiceConfig() {
			serviceNames = append(serviceNames, k)
		}
		sort.Strings(serviceNames)

		// Create reference to uhsConfig
		var uhsConfig config.Config

		// Check for input
		inputFile, err := cmd.Flags().GetString("input")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		// If input file is provided, load it
		if inputFile != "" {
			uhsConfig = config.Load(inputFile)
		} else {
			// If no input file is provided, create a new instance of the UHSConfig struct
			uhsConfig = config.Config{
				Common:   new(configCommon.Common).Default(),
				Services: make(services.ServicesConfig),
			}
		}

		// If no services are selected, prompt user to select services
		if len(selectedServices) == 0 {
			// Prompt user to select services to enable
			serviceSelectPrompt := &survey.MultiSelect{
				Message: "Select services to enable:",
				Options: serviceNames,
			}
			err := survey.AskOne(serviceSelectPrompt, &selectedServices)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}
		// Validate selected services
		if len(selectedServices) == 0 {
			fmt.Println("No services selected. Exiting...")
			os.Exit(0)
		}
		serviceListString := ""
		for _, service := range selectedServices {
			serviceListString += fmt.Sprintf("  - %v\n", service)
		}

		validateSelectionPrompt := &survey.Confirm{
			Message: fmt.Sprintf("You have selected the following services:\n%v\n Is this correct?", serviceListString),
		}
		var validateSelection bool
		err = survey.AskOne(validateSelectionPrompt, &validateSelection)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		// If selection is not valid, exit
		if !validateSelection {
			fmt.Println("Exiting...")
			os.Exit(0)

		}
		// Execute configuration for each selected service
		for _, serviceName := range selectedServices {
			fmt.Println("Configuring " + serviceName + "...")
			service := config.DefaultServiceConfig()[serviceName]
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
	configureCmd.PersistentFlags().StringP("input", "i", "", "Input file path")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configureCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configureCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
