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
	"github.com/techsquidtv/uhs-cli/models/service/servicemap"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:       "update",
	Short:     "Update your UHS instance configuration",
	Long:      `Customize and update your desired services for your UHS instance.`,
	ValidArgs: serviceNames,
	Run: func(cmd *cobra.Command, args []string) {
		// Validate files were provided
		inputFile, err := cmd.Flags().GetString("input")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		outputFile, err := cmd.Flags().GetString("output")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		var cfg config.Config
		cfg.Services = make(config.ServicesConfig)
		shared.Input(inputFile, &cfg)

		var selectedServices []string
		if len(args) > 0 {
			selectedServices = args
		}
		sort.Strings(serviceNames)

		if len(selectedServices) == 0 {
			// Prompt user to select services to update
			serviceSelectPrompt := &survey.MultiSelect{
				Message: "Select services to modify. If a selected service wasn't present in the input config, it will be added:",
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
			fmt.Println("No services selected to update. Exiting...")
			os.Exit(0)
		}

		validateSelectionPrompt := &survey.Confirm{
			Message: fmt.Sprintf("You have selected the following services:\n%v\n Is this correct?", selectedServices),
		}
		var validateSelection bool
		err = survey.AskOne(validateSelectionPrompt, &validateSelection)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		// If selection is not valid, exit
		// TODO: Make this flow better. Have user go through selection process again.
		if !validateSelection {
			fmt.Println("Exiting...")
			os.Exit(0)

		}
		// Execute configuration for each selected service
		for _, serviceName := range selectedServices {
			fmt.Println("Updating " + serviceName + "...")
			service := servicemap.Registered[serviceName]()
			service.Configure()
			cfg.Services[serviceName] = service
		}

		err = shared.Output(outputFile, &cfg)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.PersistentFlags().StringP("output", "o", "", "Output file path")
	updateCmd.PersistentFlags().StringP("input", "i", "", "Input file path")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
