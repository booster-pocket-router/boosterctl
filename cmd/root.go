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
	"os"

	"github.com/spf13/cobra"
)

var (
	host string
	port string
)

const issuer = "official cli"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "boosterctl",
	Short: "CLI client for booster",
	Long: `Allows to communicate to a booster server using its HTTP API.
Check https://github.com/booster-proj/booster to discover in detail which routes will be
involved. Through this program it is possible to list booster's sources & policies, block/unblock
specific sources and query the metrics collected using Prometheus's query language.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&host, "host-addr", "localhost", "target booster listening `host`, hostname or IP")
	rootCmd.PersistentFlags().StringVar(&port, "host-port", "7764", "target booster listening `port`")

}
