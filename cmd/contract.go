// Copyright © 2017 Weald Technology Trading
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/spf13/cobra"
)

var contractStr string
var contractAbi string

// contractCmd represents the contract command
var contractCmd = &cobra.Command{
	Use:   "contract",
	Short: "Manage contracts",
	Long:  `Call contracts directly.`,
}

func init() {
	RootCmd.AddCommand(contractCmd)
}

func contractFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&contractStr, "contract", "", "address of the contract")
	cmd.Flags().StringVar(&contractAbi, "abi", "", "ABI of the contract")
}

func parseAbi(input string) (output abi.ABI, err error) {
	var reader io.Reader
	if strings.Contains(input, string(filepath.Separator)) {
		// ABI value is a path
		reader, err = os.Open(input)
		if err != nil {
			return
		}
	} else {
		reader = strings.NewReader(input)
	}
	return abi.JSON(reader)
}
