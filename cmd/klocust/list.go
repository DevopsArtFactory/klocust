package klocust

import (
	"github.com/DevopsArtFactory/klocust/internal/klocust"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringP("namespace", "n", "", "kubernetes namespace name")
}

var listCmd = &cobra.Command{
	Aliases: []string{"ls"},
	Use:     "list",
	Short:   "Print all of Locust clusters",
	Run: func(cmd *cobra.Command, args []string) {
		namespace, err := cmd.Flags().GetString("namespace")
		if err != nil {
			log.Fatal(err)
		}

		klocust.PrintLocustDeployments(namespace)
	},
}
