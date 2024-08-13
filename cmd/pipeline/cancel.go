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
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v8"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var CancelPipelineRunCmd = &cobra.Command{
	Use:   "cancel <runId>",
	Short: "Cancel a pipeline run",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var subscriptionId, resourceGroupName, factoryName = GetArgs(cmd, args)
		var runId = args[0]

		clientFactory, err := GetClientFactory(subscriptionId)
		if err != nil {
			log.Fatalf("failed to create client: %v", err)
		}

		recursive, err := cmd.Flags().GetBool("recursive")
		if err != nil {
			log.Fatalf("failed to get the value of the flag: %v", err)
		}

		ctx := context.Background()
		_, err = clientFactory.NewPipelineRunsClient().Cancel(ctx, resourceGroupName, factoryName, runId, &armdatafactory.PipelineRunsClientCancelOptions{IsRecursive: &recursive})
		if err != nil {
			log.Fatalf("failed to finish the request: %v", err)
		} else {
			log.Print("pipeline run canceled successfully")
		}
	},
}

func init() {
	CancelPipelineRunCmd.Flags().BoolP("recursive", "r", false, "Cancel recursively")
	// triggerCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
