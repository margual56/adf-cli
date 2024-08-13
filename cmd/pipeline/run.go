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
	"encoding/json"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v8"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var RunPipelineCmd = &cobra.Command{
	Use:   "run <triggerName> [parameters]",
	Short: "Creates a run of a pipeline.",
	Long: `Creates a run of a pipeline.
For example:
  adf-cli pipeline run myPipeline "{\"param1\": \"value1\", \"param2\": \"value2\"}"
	`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var subscriptionId, resourceGroupName, factoryName = GetArgs(cmd, args)
		var pipelineName = args[0]

		var runParameters map[string]any
		if len(args) < 2 {
			runParameters = make(map[string]any)
		} else {
			// Parse the parameters as J
			err := json.Unmarshal([]byte(args[1]), &runParameters)
			if err != nil {
				log.Fatalf("Unable to marshal JSON due to %s", err)
			}
		}

		clientFactory, err := GetClientFactory(subscriptionId)
		if err != nil {
			log.Fatalf("failed to create client: %v", err)
		}

		ctx := context.Background()
		log.Println("running pipeline...")
		res, err := clientFactory.NewPipelinesClient().CreateRun(ctx, resourceGroupName, factoryName, pipelineName, &armdatafactory.PipelinesClientCreateRunOptions{ReferencePipelineRunID: nil,
			IsRecovery:        nil,
			StartActivityName: nil,
			StartFromFailure:  nil,
			Parameters:        runParameters,
		})
		if err != nil {
			log.Fatalf("failed to finish the request: %v", err)
		} else {
			log.Printf("success! Run ID: %v\n", *res.RunID)
		}
	},
}

func init() {
	// triggerCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
