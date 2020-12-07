package klocust

import (
	"context"
	"github.com/DevopsArtFactory/klocust/internal/klocust"
	"github.com/spf13/cobra"
	"io"
)

func doDelete(_ context.Context, _ io.Writer, _ *cobra.Command, args []string) error {
	locustName := args[0]
	return klocust.DeleteLocust(opts.Namespace, locustName)
}

func NewDeleteCmd() *cobra.Command {
	return NewCmd("delete").
		WithDescription("Delete klocust cluster").
		WithCommonFlags().
		ExactArgs(1, doDelete)
}
