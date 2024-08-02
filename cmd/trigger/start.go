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

package trigger

import (
	"context"
	"log"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var StartTriggerCmd = &cobra.Command{
	Use:   "start <triggerName>",
	Short: "Start a trigger in a data factory.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var subscriptionId, resourceGroupName, factoryName = GetArgs(cmd, args)
		var triggerName = args[0]

		clientFactory, err := GetClientFactory(subscriptionId)
		if err != nil {
			log.Fatalf("failed to create client: %v", err)
		}

		ctx := context.Background()
		log.Printf("starting trigger %q", triggerName)
		poller, err := clientFactory.NewTriggersClient().BeginStart(ctx, resourceGroupName, factoryName, triggerName, nil)
		if err != nil {
			log.Fatalf("failed to finish the request: %v", err)
		}
		_, err = poller.PollUntilDone(ctx, nil)
		if err != nil {
			log.Fatalf("failed to pull the result: %v", err)
		}

		log.Printf("trigger %q started successfully", triggerName)
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
