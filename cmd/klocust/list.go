package klocust

import (
	"context"
	"github.com/DevopsArtFactory/klocust/cmd/builder"
	"github.com/DevopsArtFactory/klocust/internal/klocust"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"io"
)

var (
	namespace string
)

func list(_ context.Context, _ io.Writer) error {
	return klocust.PrintLocustDeployments(namespace)
}

func NewCmdList() *cobra.Command {
	return builder.NewCmd("list").
		WithDescription("Print all of Locust clusters").
		WithFlags(func(f *pflag.FlagSet) {
			f.StringVarP(&namespace, "namespace", "n", "", "Kubernetes namespace")
		}).NoArgs(list)
}
