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

package pipeline

import (
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v8"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Get the subscriptionId, resourceGroupName, and factoryName from the environment variables or command line flags.
// If the values are not set, the program will exit.
func GetArgs(cmd *cobra.Command, args []string) (string, string, string) {
	var subscriptionId string
	if viper.IsSet("subscriptionId") {
		subscriptionId = viper.GetString("subscriptionId")
	} else if cmd.Flag("subscriptionId").Value.String() != "" {
		subscriptionId = cmd.Flag("subscriptionId").Value.String()
	} else {
		log.Fatalf("subscriptionId is required")
	}

	var resourceGroupName string
	if viper.IsSet("resourceGroupName") {
		resourceGroupName = viper.GetString("resourceGroupName")
	} else if cmd.Flag("resourceGroupName").Value.String() != "" {
		resourceGroupName = cmd.Flag("resourceGroupName").Value.String()
	} else {
		log.Fatalf("resourceGroupName is required")
	}

	var factoryName string
	if viper.IsSet("factoryName") {
		factoryName = viper.GetString("factoryName")
	} else if cmd.Flag("factoryName").Value.String() != "" {
		factoryName = cmd.Flag("factoryName").Value.String()
	} else {
		log.Fatalf("factoryName is required")
	}

	return subscriptionId, resourceGroupName, factoryName
}

// Create a new client factory with the default Azure credential
// If the credential cannot be obtained, the program will exit.
func GetClientFactory(subscriptionId string) (*armdatafactory.ClientFactory, error) {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	return armdatafactory.NewClientFactory(subscriptionId, cred, nil)
}

// triggerCmd represents the trigger command
var PipelineCmd = &cobra.Command{
	Use: "pipeline [command]",
	Run: func(cmd *cobra.Command, args []string) {
		// Explain that trigger needs a subcommand and suggest the subcommands available
		fmt.Println("param needs a subcommand. Available subcommands are:")
		for _, cmd := range cmd.Commands() {
			fmt.Printf("  %s\n", cmd.Use)
		}
	},
}

func init() {
	PipelineCmd.AddCommand(RunPipelineCmd)
	PipelineCmd.AddCommand(ListPipelineCmd)

	// rootCmd.AddCommand(triggerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// triggerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// triggerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
