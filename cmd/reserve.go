// Copyright © 2019 KIM KeepInMind GmbH/srl
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package cmd

import (
	"fmt"
	"io"
	"net"
	"os"

	"github.com/booster-proj/boosterctl/client"
	"github.com/spf13/cobra"
)

// reserveCmd represents the reserve command
var reserveCmd = &cobra.Command{
	Use:   "reserve `source_id` `hosts...`",
	Short: "Make booster reserve `source_id` only for connections to `hosts`",
	Long: `Perform an HTTP request to "/reserve.json", making booster reserve
source "source_id" only for connections to the list "hosts", otherwise the other sources
are used.`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		cl, err := client.New(net.JoinHostPort(host, port))
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		status, r, err := cl.AddPolicy("reserve", struct{
			*client.PolicyReq
			Hosts []string
		}{
			PolicyReq: &client.PolicyReq{
				SourceID: args[0],
				Issuer:   issuer,
			},
			Hosts: args[1:],
		})
		fmt.Printf("Status: %v\n", status)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		io.Copy(os.Stderr, r)
	},
}

func init() {
	policiesCmd.AddCommand(reserveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// reserveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// reserveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
