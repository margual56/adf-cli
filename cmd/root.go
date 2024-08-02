/*
Copyright © 2024 Marcos Gutiérrez Alonso marcos56@mailbox.org

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/margual56/adf-cli/cmd/param"
	"github.com/margual56/adf-cli/cmd/trigger"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "adf-cli",
	Short: "A simple CLI to interact with Azure Data Factory.",
	Long:  `With this CLI you can manage triggers from a factory in Azure Data Factory. More features will be added in the future.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(trigger.TriggerCmd)
	rootCmd.AddCommand(param.ParamCmd)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.adf-cli.yaml)")

	rootCmd.PersistentFlags().String("subscriptionId", "", "The subscription identifier.")
	rootCmd.PersistentFlags().String("resourceGroupName", "", "The resource group name.")
	rootCmd.PersistentFlags().String("factoryName", "", "The factory name.")

	// rootCmd.MarkPersistentFlagRequired("subscriptionId")
	// rootCmd.MarkPersistentFlagRequired("resourceGroupName")
	// rootCmd.MarkPersistentFlagRequired("factoryName")

	viper.BindEnv("subscriptionId", "SUBSCRIPTION_ID")
	viper.BindEnv("resourceGroupName", "RESOURCE_GROUP_NAME")
	viper.BindEnv("factoryName", "FACTORY_NAME")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
