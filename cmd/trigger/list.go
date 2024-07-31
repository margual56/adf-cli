package trigger

import (
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var ListTriggerCmd = &cobra.Command{
	Use:   "list",
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
