package klocust

import (
	"context"
	"github.com/DevopsArtFactory/klocust/internal/klocust"
	"github.com/spf13/cobra"
	"io"
)

func doApply(_ context.Context, _ io.Writer, _ *cobra.Command, args []string) error {
	locustName := args[0]
	return klocust.ApplyLocust(opts.Namespace, locustName)
}

func NewApplyCMD() *cobra.Command {
	return NewCmd("apply").
		WithDescription("Apply klocust cluster.\n"+
			"This cluster will be created if it doesn't exist yet.").
		WithCommonFlags().
		ExactArgs(1, doApply)
}
