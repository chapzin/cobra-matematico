package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "soma",
	Short: "Essa tag você soma uma quantidade infinita de numeros",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		addInt(args)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func addInt(args []string) {
	var sum int
	for _, ival := range args {
		itemp, err := strconv.Atoi(ival)
		if err != nil {
			fmt.Println(err)
		}
		sum = sum + itemp
	}

	fmt.Printf("A soma dos numeros %s adicionado é %d \n", args, sum)
}
