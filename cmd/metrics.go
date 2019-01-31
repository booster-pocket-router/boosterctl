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
	"net/url"
	"os"

	"github.com/booster-proj/booster.cli/client"
	"github.com/spf13/cobra"
)

// metricsCmd represents the metrics command
var metricsCmd = &cobra.Command{
	Use:   "metrics `query`",
	Short: "Execute `query` on the local Prometheus server, if available",
	Long: `Perform an HTTP request to "/metrics.json" with the provided query. The request
is forwarded to the local Prometheus server instance if available. Enclose the query into quotes
if it contains spaces.
Some query examples:
  - sum by (source) (booster_network_receive_bytes) # Count of received bytes
  - sum by (source) (rate(booster_network_receive_bytes{}[10s])*8) # Download bandwidth

Note that these are executed as instant queries, i.e. their results refer to the current instant.
For more: https://prometheus.io/docs/prometheus/latest/querying/api/#instant-queries`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cl, err := client.New(net.JoinHostPort(host, port))
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		val := url.Values{}
		val.Set("query", args[0])
		status, r, err := cl.QueryMetrics(&val)
		fmt.Printf("Status: %v\n", status)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		io.Copy(os.Stderr, r)
		// this response does not end with \n
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(metricsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// metricsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// metricsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
