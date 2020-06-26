package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// evenCmd represents the even command
var evenCmd = &cobra.Command{
	Use:   "pares",
	Short: "Comando soma apenas os numeros pares",
	Long:  `o objetivo desse comando é somar apenas os numeros pares`,
	Run: func(cmd *cobra.Command, args []string) {
		var evenSum int
		for _, ival := range args {
			itemp, _ := strconv.Atoi(ival)
			if itemp%2 == 0 {
				evenSum = evenSum + itemp
			}
		}
		fmt.Printf("Voce informou os numeros %s a soma apenas dos pares é %d\n", args, evenSum)
	},
}

func init() {
	addCmd.AddCommand(evenCmd)

}
