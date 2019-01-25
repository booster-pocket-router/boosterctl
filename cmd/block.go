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
	"net/http"

	"github.com/booster-proj/booster.cli/client"
	"github.com/spf13/cobra"
)

// blockCmd represents the block command
var blockCmd = &cobra.Command{
	Use:   "block",
	Short: "Make booster block the sources specified",
	Long: `Perform an HTTP request to "/block.json" for each source specified, making
booster add a block policy on it, i.e. the source will no longer be used.
Outputs the errors returned if any.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, v := range args {
			block(v)
		}
	},
}

func block(source string) {
	cl, err := client.New(net.JoinHostPort(host, port))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	resp, err := cl.Post("/sources/"+source+"/block.json", nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if resp.StatusCode != http.StatusCreated {
		fmt.Printf("Error: Response Status: %v\n", resp.Status)
		return
	}

	defer resp.Body.Close()
	fmt.Printf("Blocked: %s\n", source)
}

func init() {
	rootCmd.AddCommand(blockCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// blockCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// blockCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
