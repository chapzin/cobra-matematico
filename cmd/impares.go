/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// oddCmd represents the odd command
var oddCmd = &cobra.Command{
	Use:   "impares",
	Short: "Esse comando soma apenas os numeros impares",
	Long:  `Comando soma apenas numeros impares`,
	Run: func(cmd *cobra.Command, args []string) {
		var oddSum int
		for _, ival := range args {
			itemp, _ := strconv.Atoi(ival)
			if itemp%2 != 0 {
				oddSum = oddSum + itemp
			}
		}
		fmt.Printf("Voce informou os numeros %s a soma apenas dos impares é %d\n", args, oddSum)
	},
}

func init() {
	addCmd.AddCommand(oddCmd)

}
