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
	"net"
	"os"

	"github.com/booster-proj/booster.cli/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	host string
	port string
	cl   *client.Client
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "booster.cli",
	Short: "CLI client for booster",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&host, "api.host", "localhost", "target booster listening `host`, hostname or IP")
	rootCmd.PersistentFlags().StringVar(&port, "api.port", "7764", "target booster listening `port`")

	initClient() // has to run only once
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.AutomaticEnv() // read in environment variables that match
}

func initClient() {
	var err error

	cl, err = client.New(net.JoinHostPort(host, port))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
