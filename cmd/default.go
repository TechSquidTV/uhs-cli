/*
Copyright © 2023 TechSquidTV
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/techsquidtv/uhs-cli/cmd/shared"
	"github.com/techsquidtv/uhs-cli/models/config"
	"github.com/techsquidtv/uhs-cli/models/global"
	"github.com/techsquidtv/uhs-cli/models/service/servicemap"
)

// defaultCmd represents the default command
var defaultCmd = &cobra.Command{
	Use:   "default",
	Short: "Get the default configuration for UHS",
	Long: `Get the default configuration for UHS. 
		This will output the default configuration file, meant to be overwritten manually of via the configure command.`,
	Run: func(cmd *cobra.Command, args []string) {
		uhsConfig := config.Config{
			Global:   global.NewGlobal(),
			Services: make(config.ServicesConfig),
		}

		// Set services to default
		for k, v := range servicemap.Registered {
			uhsConfig.Services[k] = v()
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
	rootCmd.AddCommand(defaultCmd)
	defaultCmd.PersistentFlags().StringP("output", "o", "", "Output file path")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// defaultCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// defaultCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
