package klocust

import (
	"fmt"
	"github.com/DevopsArtFactory/klocust/internal/klocust"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/api/errors"
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

		if err := klocust.PrintLocustDeployments(namespace); err != nil {
			if errors.IsUnauthorized(err) {
				fmt.Println("Check your kubeconfig: ", err)
			}
			log.Fatal(err)
		}
	},
}
