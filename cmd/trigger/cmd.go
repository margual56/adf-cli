/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package trigger

import (
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v8"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func GetArgs(cmd *cobra.Command, args []string) (string, string, string) {
	var subscriptionId string
	if viper.IsSet("subscriptionId") {
		subscriptionId = viper.GetString("subscriptionId")
	} else {
		subscriptionId = cmd.Flag("subscriptionId").Value.String()
	}

	var resourceGroupName string
	if viper.IsSet("resourceGroupName") {
		resourceGroupName = viper.GetString("resourceGroupName")
	} else {
		resourceGroupName = cmd.Flag("resourceGroupName").Value.String()
	}

	var factoryName string
	if viper.IsSet("factoryName") {
		factoryName = viper.GetString("factoryName")
	} else {
		factoryName = cmd.Flag("factoryName").Value.String()
	}

	return subscriptionId, resourceGroupName, factoryName
}

func GetClientFactory(subscriptionId string) (*armdatafactory.ClientFactory, error) {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	return armdatafactory.NewClientFactory(subscriptionId, cred, nil)
}

// triggerCmd represents the trigger command
var TriggerCmd = &cobra.Command{
	Use:   "trigger",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("trigger called")
	},
}

func init() {
	TriggerCmd.AddCommand(GetTriggerCmd)
	TriggerCmd.AddCommand(ListTriggerCmd)
	TriggerCmd.AddCommand(StartTriggerCmd)
	TriggerCmd.AddCommand(StopTriggerCmd)

	// rootCmd.AddCommand(triggerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// triggerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// triggerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
