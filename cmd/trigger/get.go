/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package trigger

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v8"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var GetTriggerCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("client: making http request to server...\n")

		subscriptionId := cmd.Flag("subscriptionId").Value.String()
		resourceGroupName := cmd.Flag("resourceGroupName").Value.String()
		factoryName := cmd.Flag("factoryName").Value.String()
		triggerName := cmd.Flag("triggerName").Value.String()

		cred, err := azidentity.NewDefaultAzureCredential(nil)
		if err != nil {
			log.Fatalf("failed to obtain a credential: %v", err)
		}
		ctx := context.Background()
		clientFactory, err := armdatafactory.NewClientFactory(subscriptionId, cred, nil)
		if err != nil {
			log.Fatalf("failed to create client: %v", err)
		}
		res, err := clientFactory.NewTriggersClient().Get(ctx, resourceGroupName, factoryName, triggerName, &armdatafactory.TriggersClientGetOptions{IfNoneMatch: nil})
		if err != nil {
			log.Fatalf("failed to finish the request: %v", err)
		}

		fmt.Printf("client: received response from server: %v\n", res.Properties)
	},
}

func init() {
	// triggerCmd.AddCommand(getCmd)
	GetTriggerCmd.Flags().String("triggerName", "", "The trigger name.")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
