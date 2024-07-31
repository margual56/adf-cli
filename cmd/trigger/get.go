/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package trigger

import (
	"fmt"
	"net/http"
	"os"

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

		requestUrl := fmt.Sprint("https://management.azure.com/subscriptions/", subscriptionId, "/resourceGroups/", resourceGroupName, "/providers/Microsoft.DataFactory/factories/", factoryName, "/triggers/", triggerName, "?api-version=2018-06-01")

		res, err := http.Get(requestUrl)
		if err != nil {
			fmt.Printf("error making http request: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("client: got response!\n")
		fmt.Printf("client: status code: %d\n", res.StatusCode)
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
