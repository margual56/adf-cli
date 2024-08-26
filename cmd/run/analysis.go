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

package run

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v8"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var RunAnalysisCmd = &cobra.Command{
	Use:   "analysis",
	Short: "Display the properties of a trigger by name.",
	Run: func(cmd *cobra.Command, args []string) {
		var subscriptionId, resourceGroupName, factoryName = GetArgs(cmd, args)

		clientFactory, err := GetClientFactory(subscriptionId)
		if err != nil {
			log.Fatalf("failed to create client: %v", err)
		}

		ctx := context.Background()

		// Create the filters for the query from the command line arguments
		filters := []*armdatafactory.RunQueryFilter{}
		pipelineName, _ := cmd.Flags().GetString("pipelineName")
		if pipelineName != "" {
			filters = append(filters, &armdatafactory.RunQueryFilter{
				Operand:  to.Ptr(armdatafactory.RunQueryFilterOperandPipelineName),
				Operator: to.Ptr(armdatafactory.RunQueryFilterOperatorEquals),
				Values:   []*string{to.Ptr(pipelineName)},
			})
		}

		allRuns := []*armdatafactory.PipelineRun{}

		// Query the runs
		res, err := clientFactory.NewPipelineRunsClient().QueryByFactory(ctx, resourceGroupName, factoryName, armdatafactory.RunFilterParameters{
			Filters:           filters,
			LastUpdatedAfter:  to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-06-01T00:00:00.000Z"); return t }()),
			LastUpdatedBefore: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-01T00:00:00.000Z"); return t }()),
			OrderBy: []*armdatafactory.RunQueryOrderBy{
				{
					Order:   to.Ptr(armdatafactory.RunQueryOrderASC),
					OrderBy: to.Ptr(armdatafactory.RunQueryOrderByFieldRunStart),
				},
			},
		}, nil)
		if err != nil {
			log.Fatalf("failed to finish the request: %v", err)
		}

		allRuns = append(allRuns, res.Value...)
		for res.ContinuationToken != nil {
			res, err = clientFactory.NewPipelineRunsClient().QueryByFactory(ctx, resourceGroupName, factoryName, armdatafactory.RunFilterParameters{
				Filters:           filters,
				LastUpdatedAfter:  to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-06-01T00:00:00.000Z"); return t }()),
				LastUpdatedBefore: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-01T00:00:00.000Z"); return t }()),
				ContinuationToken: res.ContinuationToken,
				OrderBy: []*armdatafactory.RunQueryOrderBy{
					{
						Order:   to.Ptr(armdatafactory.RunQueryOrderASC),
						OrderBy: to.Ptr(armdatafactory.RunQueryOrderByFieldRunStart),
					},
				},
			}, nil)
			allRuns = append(allRuns, res.Value...)
		}

		log.Printf("Got %v runs", len(allRuns))

		var runs []*armdatafactory.PipelineRun

		// Filter the runs that started after the specified time
		startAfterStr, _ := cmd.Flags().GetString("startAfter")
		endBeforeStr, _ := cmd.Flags().GetString("endBefore")

		if startAfterStr != "" && endBeforeStr != "" {
			startAfter, err := time.Parse(time.RFC3339, startAfterStr)
			if err != nil {
				log.Fatalf("failed to parse the startAfter time: %v", err)
			}
			endBefore, err := time.Parse(time.RFC3339, endBeforeStr)
			if err != nil {
				log.Fatalf("failed to parse the endBefore time: %v", err)
			}

			for _, run := range allRuns {
				if run.RunStart.After(startAfter) && run.RunEnd.Before(endBefore) {
					runs = append(runs, run)
				}
			}
		} else if startAfterStr != "" {
			startAfter, err := time.Parse(time.RFC3339, startAfterStr)
			if err != nil {
				log.Fatalf("failed to parse the startAfter time: %v", err)
			}

			for _, run := range allRuns {
				if run.RunStart.After(startAfter) {
					runs = append(runs, run)
				}
			}
		} else if endBeforeStr != "" {
			endBefore, err := time.Parse(time.RFC3339, endBeforeStr)
			if err != nil {
				log.Fatalf("failed to parse the endBefore time: %v", err)
			}

			for _, run := range allRuns {
				if run.RunEnd.Before(endBefore) {
					runs = append(runs, run)
				}
			}
		} else {
			runs = res.Value
		}

		log.Printf("Got %v runs after filtering", len(runs))

		var totalRuntime int32 = 0
		// Print the runs
		fmt.Print("Pipeline name;Start time;End time;Duration (ms);Run parameters;Cost;Reason for run;Run link;Status;Notes;Manual notes\n")
		for _, run := range runs {
			runLink := fmt.Sprintf("https://adf.azure.com/en/monitoring/pipelineruns/%s?factory=%%2Fsubscriptions%%2F%s%%2FresourceGroups%%2F%s%%2Fproviders%%2FMicrosoft.DataFactory%%2Ffactories%%2F%s", *run.RunID, subscriptionId, resourceGroupName, factoryName)
			parameters := ""

			for k, v := range run.Parameters {
				parameters += fmt.Sprintf("%s=%s;", k, *v)
			}

			fmt.Printf("%s;%v;%v;%v;\"%v\";;\"%s\";%s;%s;\"%s\";\n", *run.PipelineName, *run.RunStart, *run.RunEnd, *run.DurationInMs, parameters, *run.InvokedBy.Name, runLink, *run.Status, *run.Message)

			totalRuntime += *run.DurationInMs
		}

		log.Printf("Total runtime: %vms", totalRuntime)
	},
}

func init() {
	// triggerCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")
	RunAnalysisCmd.PersistentFlags().StringP("pipelineName", "n", "", "Only query runs with the specified pipeline name.")
	RunAnalysisCmd.PersistentFlags().StringP("startAfter", "s", "", "Only query runs that started after the specified time.")
	RunAnalysisCmd.PersistentFlags().StringP("endBefore", "e", "", "Only query runs that ended before the specified time.")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
