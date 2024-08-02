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

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var GetParamCmd = &cobra.Command{
	Use:   "get <globalParameterName>",
	Short: "Get a global parameter by name from a data factory.",
	Run: func(cmd *cobra.Command, args []string) {
		var subscriptionId, resourceGroupName, factoryName = GetArgs(cmd, args)
		parameterGroupName := cmd.Flag("group").Value.String()
		globalParameterName := args[0]

		clientFactory, err := GetClientFactory(subscriptionId)
		if err != nil {
			log.Fatalf("failed to create client: %v", err)
		}

		ctx := context.Background()
		res, err := clientFactory.NewGlobalParametersClient().Get(ctx, resourceGroupName, factoryName, parameterGroupName, nil)
		if err != nil {
			log.Fatalf("failed to finish the request: %v", err)
		}

		if val, ok := res.Properties[globalParameterName]; !ok {
			log.Fatalf("Parameter %q not found in group %q", globalParameterName, parameterGroupName)
		} else {
			log.Printf("Parameter %q has type %s with value %q", globalParameterName, *val.Type, val.Value)
		}
	},
}

func init() {
	GetParamCmd.Flags().StringP("group", "g", "default", "Group name where the global parameter is stored.")
	// triggerCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
