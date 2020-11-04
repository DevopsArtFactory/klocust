package klocust

import (
	"context"
	"github.com/DevopsArtFactory/klocust/internal/klocust"
	"github.com/spf13/cobra"
	"io"
)

func doInit(_ context.Context, _ io.Writer, _ *cobra.Command, args []string) error {
	kLocustName := args[0]
	return klocust.InitLocust(opts.Namespace, kLocustName)
}

func NewInitCmd() *cobra.Command {
	return NewCmd("init").
		WithDescription("init").
		WithCommonFlags().
		ExactArgs(1, doInit)
}
