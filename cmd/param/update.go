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
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v8"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var UpdateParamCmd = &cobra.Command{
	Use:   "update <globalParameterName>",
	Short: "Create or update a global parameter in a data factory",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var subscriptionId, resourceGroupName, factoryName = GetArgs(cmd, args)

		var paramType *armdatafactory.GlobalParameterType
		switch tmp_type := cmd.Flag("type").Value.String(); tmp_type {
		case "int":
			paramType = to.Ptr(armdatafactory.GlobalParameterTypeInt)
		case "string":
			paramType = to.Ptr(armdatafactory.GlobalParameterTypeString)
		case "bool":
			paramType = to.Ptr(armdatafactory.GlobalParameterTypeBool)
		case "array":
			paramType = to.Ptr(armdatafactory.GlobalParameterTypeArray)
		default:
			log.Fatalf("Invalid parameter type")
		}

		paramValue := cmd.Flag("value").Value.String()
		paramGroup := cmd.Flag("group").Value.String()

		var globalParameterName = args[0]

		clientFactory, err := GetClientFactory(subscriptionId)
		if err != nil {
			log.Fatalf("failed to create client: %v", err)
		}

		ctx := context.Background()
		res, err := clientFactory.NewGlobalParametersClient().CreateOrUpdate(ctx, resourceGroupName, factoryName, paramGroup, armdatafactory.GlobalParameterResource{
			Properties: map[string]*armdatafactory.GlobalParameterSpecification{
				globalParameterName: {
					Type:  paramType,
					Value: paramValue,
				},
			},
		}, nil)
		if err != nil {
			log.Fatalf("failed to finish the request: %v", err)
		}
		// You could use response here. We use blank identifier for just demo purposes.
		_ = res
		// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// res.GlobalParameterResource = armdatafactory.GlobalParameterResource{
		// 	Name: to.Ptr("default"),
		// 	Type: to.Ptr("Microsoft.DataFactory/factories/globalParameters"),
		// 	Etag: to.Ptr("0a008ad4-0000-0000-0000-5b245c6e0000"),
		// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/globalParameters/default"),
		// 	Properties: map[string]*armdatafactory.GlobalParameterSpecification{
		// 		"waitTime": &armdatafactory.GlobalParameterSpecification{
		// 			Type: to.Ptr(armdatafactory.GlobalParameterTypeInt),
		// 			Value: float64(5),
		// 		},
		// 	},
		// }

		log.Printf("Variable %q from group %q set to %q successfully", globalParameterName, paramGroup, paramValue)
	},
}

func init() {
	UpdateParamCmd.Flags().StringP("group", "g", "default", "The group the parameter belongs to")

	UpdateParamCmd.Flags().StringP("type", "t", "string", "The type of the parameter")
	UpdateParamCmd.MarkFlagRequired("type")

	UpdateParamCmd.Flags().StringP("value", "v", "", "The value of the parameter")
	UpdateParamCmd.MarkFlagRequired("value")

	// triggerCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
