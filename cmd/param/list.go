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

package param

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var ListParamCmd = &cobra.Command{
	Use:   "list <globalParameterName>",
	Short: "List all global parameters from a factory",
	Run: func(cmd *cobra.Command, args []string) {
		var subscriptionId, resourceGroupName, factoryName = GetArgs(cmd, args)

		clientFactory, err := GetClientFactory(subscriptionId)
		if err != nil {
			log.Fatalf("failed to create client: %v", err)
		}

		ctx := context.Background()
		pager := clientFactory.NewGlobalParametersClient().NewListByFactoryPager(resourceGroupName, factoryName, nil)
		for pager.More() {
			page, err := pager.NextPage(ctx)
			if err != nil {
				log.Fatalf("failed to advance page: %v", err)
			}
			for _, v := range page.Value {
				fmt.Printf("%q {\n", *v.Name)

				// Print the properties of the global parameter
				for k, v := range v.Properties {
					fmt.Printf("    %q: %q,\n", k, v.Value)
				}

				fmt.Println("}")
			}
			// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
			// page.GlobalParameterListResponse = armdatafactory.GlobalParameterListResponse{
			// 	Value: []*armdatafactory.GlobalParameterResource{
			// 		{
			// 			Name: to.Ptr("default"),
			// 			Type: to.Ptr("Microsoft.DataFactory/factories/globalParameters"),
			// 			Etag: to.Ptr("da00a1c3-0000-0400-0000-6241f3290000"),
			// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/globalParameters/default"),
			// 			Properties: map[string]*armdatafactory.GlobalParameterSpecification{
			// 				"copyPipelineTest": &armdatafactory.GlobalParameterSpecification{
			// 					Type: to.Ptr(armdatafactory.GlobalParameterTypeObject),
			// 					Value: map[string]any{
			// 						"mySinkDatasetFolderPath": "exampleOutput",
			// 						"mySourceDatasetFolderPath": "exampleInput/",
			// 						"testingEmptyContextParams": "",
			// 					},
			// 				},
			// 				"mySourceDatasetFolderPath": &armdatafactory.GlobalParameterSpecification{
			// 					Type: to.Ptr(armdatafactory.GlobalParameterTypeString),
			// 					Value: "input",
			// 				},
			// 				"url": &armdatafactory.GlobalParameterSpecification{
			// 					Type: to.Ptr(armdatafactory.GlobalParameterTypeString),
			// 					Value: "https://testuri.test",
			// 				},
			// 				"validADFOffice365Uris": &armdatafactory.GlobalParameterSpecification{
			// 					Type: to.Ptr(armdatafactory.GlobalParameterTypeArray),
			// 					Value: []any{
			// 						"https://testuri.test",
			// 						"https://testuri.test",
			// 					},
			// 				},
			// 				"waitTime": &armdatafactory.GlobalParameterSpecification{
			// 					Type: to.Ptr(armdatafactory.GlobalParameterTypeInt),
			// 					Value: float64(5),
			// 				},
			// 			},
			// 	}},
			// }
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
