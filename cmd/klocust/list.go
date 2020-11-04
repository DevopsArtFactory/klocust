package klocust

import (
	"context"
	"github.com/DevopsArtFactory/klocust/internal/klocust"
	"github.com/spf13/cobra"
	"io"
)

func doList(_ context.Context, _ io.Writer) error {
	return klocust.PrintLocustDeployments(opts.Namespace)
}

func NewCmdList() *cobra.Command {
	return NewCmd("list").
		WithDescription("Display all of Locust clusters").
		WithCommonFlags().
		NoArgs(doList)
}
