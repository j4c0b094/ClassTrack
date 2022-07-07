package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add task",
	Short: "add task",
	Run: func(cmd *cobra.Command, args []string) {
		add()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func add() {
	fmt.Println("hello")
}
