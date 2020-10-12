package klocust

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Print all of Locust clusters",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list result...")
	},
}
