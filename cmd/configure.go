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
	"github.com/techsquidtv/uhs-cli/cmd/shared"
	"github.com/techsquidtv/uhs-cli/models/config"
	"github.com/techsquidtv/uhs-cli/models/global"
	"github.com/techsquidtv/uhs-cli/models/service/servicemap"
)

// Return each key from the default config map
var serviceNames = servicemap.Keys(servicemap.Registered)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:       "configure",
	Short:     "Configure your UHS instance",
	Long:      `Customize and configure your desired services for your UHS instance.`,
	ValidArgs: serviceNames,
	Run: func(cmd *cobra.Command, args []string) {
		// Create a new instance of the UHSConfig struct
		var selectedServices []string
		if len(args) > 0 {
			selectedServices = args
		}
		uhsConfig := config.Config{
			Global:   global.NewGlobal(),
			Services: make(config.ServicesConfig),
		}
		sort.Strings(serviceNames)

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
		err := survey.AskOne(validateSelectionPrompt, &validateSelection)
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
			service := servicemap.Registered[serviceName]()
			service.Configure()
			uhsConfig.Services[serviceName] = service
		}
		// Output
		outputFile, err := cmd.Flags().GetString("output")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		err = shared.Output(outputFile, &uhsConfig)
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
