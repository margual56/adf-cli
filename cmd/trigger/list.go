package trigger

import (
	"context"
	"fmt"
	"log"

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
		var subscriptionId, resourceGroupName, factoryName = GetArgs(cmd, args)

		clientFactory, err := GetClientFactory(subscriptionId)
		if err != nil {
			log.Fatalf("failed to create client: %v", err)
		}
		pager := clientFactory.NewTriggersClient().NewListByFactoryPager(resourceGroupName, factoryName, nil)

		ctx := context.Background()
		for pager.More() {
			page, err := pager.NextPage(ctx)
			if err != nil {
				log.Fatalf("failed to advance page: %v", err)
			}
			for _, v := range page.Value {
				fmt.Println(*v.Name)
			}
			// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
			// page.TriggerListResponse = armdatafactory.TriggerListResponse{
			// 	Value: []*armdatafactory.TriggerResource{
			// 		{
			// 			Name: to.Ptr("exampleTrigger"),
			// 			Type: to.Ptr("Microsoft.DataFactory/factories/triggers"),
			// 			Etag: to.Ptr("0a008ed4-0000-0000-0000-5b245c740000"),
			// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/triggers/exampleTrigger"),
			// 			Properties: &armdatafactory.ScheduleTrigger{
			// 				Type: to.Ptr("ScheduleTrigger"),
			// 				Description: to.Ptr("Example description"),
			// 				RuntimeState: to.Ptr(armdatafactory.TriggerRuntimeStateStarted),
			// 				Pipelines: []*armdatafactory.TriggerPipelineReference{
			// 					{
			// 						Parameters: map[string]any{
			// 							"OutputBlobNameList": []any{
			// 								"exampleoutput.csv",
			// 							},
			// 						},
			// 						PipelineReference: &armdatafactory.PipelineReference{
			// 							Type: to.Ptr(armdatafactory.PipelineReferenceTypePipelineReference),
			// 							ReferenceName: to.Ptr("examplePipeline"),
			// 						},
			// 				}},
			// 				TypeProperties: &armdatafactory.ScheduleTriggerTypeProperties{
			// 					Recurrence: &armdatafactory.ScheduleTriggerRecurrence{
			// 						EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:55:14.905Z"); return t}()),
			// 						Frequency: to.Ptr(armdatafactory.RecurrenceFrequencyMinute),
			// 						Interval: to.Ptr[int32](4),
			// 						StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:39:14.905Z"); return t}()),
			// 						TimeZone: to.Ptr("UTC"),
			// 					},
			// 				},
			// 			},
			// 	}},
			// }
		}
	},
}
