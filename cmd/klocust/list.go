package klocust

import (
	"fmt"
	"github.com/DevopsArtFactory/klocust/internal/klocust"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/api/errors"
	"log"
)

func list(cmd *cobra.Command, args []string) {
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
}

func NewCmdList() *cobra.Command {
	newCmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Args: func(cmd *cobra.Command, args []string) error {
			return cobra.OnlyValidArgs(cmd, args)
		},
		ValidArgs: []string{"list", "ls"},
		Short:     "Print all of Locust clusters",
		Run:       list,
	}

	newCmd.Flags().StringP("namespace", "n", "", "kubernetes namespace name")
	return newCmd
}
