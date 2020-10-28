// Builder from https://github.com/GoogleContainerTools/skaffold
package builder

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"io"
	"k8s.io/apimachinery/pkg/api/errors"
)

// Builder is used to build cobra commands.
type Builder interface {
	WithDescription(description string) Builder
	WithLongDescription(long string) Builder
	WithExample(comment, command string) Builder
	WithFlags(adder func(*pflag.FlagSet)) Builder
	WithCommonFlags() Builder
	Hidden() Builder
	ExactArgs(argCount int, action func(context.Context, io.Writer, *cobra.Command, []string) error) *cobra.Command
	NoArgs(action func(context.Context, io.Writer) error) *cobra.Command
}

type builder struct {
	cmd cobra.Command
}

// NewCmd creates a new command builder.
func NewCmd(use string) *builder {
	return &builder{
		cmd: cobra.Command{
			Use: use,
		},
	}
}

func (b *builder) WithDescription(description string) Builder {
	b.cmd.Short = description
	return b
}

func (b *builder) WithLongDescription(long string) Builder {
	b.cmd.Long = long
	return b
}

func (b *builder) WithExample(comment, command string) Builder {
	if b.cmd.Example != "" {
		b.cmd.Example += "\n"
	}
	b.cmd.Example += fmt.Sprintf("  # %s\n  klocust %s\n", comment, command)
	return b
}

func (b *builder) WithCommonFlags() Builder {
	AddFlags(&b.cmd)
	return b
}

func (b *builder) WithFlags(adder func(*pflag.FlagSet)) Builder {
	adder(b.cmd.Flags())
	return b
}

func (b *builder) Hidden() Builder {
	b.cmd.Hidden = true
	return b
}

func (b *builder) SetAliases(alias []string) Builder {
	b.cmd.Aliases = alias
	return b
}

func (b *builder) ExactArgs(argCount int, action func(context.Context, io.Writer, *cobra.Command, []string) error) *cobra.Command {
	b.cmd.Args = cobra.ExactArgs(argCount)
	b.cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return handleWellKnownErrors(action(b.cmd.Context(), b.cmd.OutOrStdout(), cmd, args))
	}
	return &b.cmd
}

func (b *builder) NoArgs(action func(context.Context, io.Writer) error) *cobra.Command {
	b.cmd.Args = cobra.NoArgs
	b.cmd.RunE = func(*cobra.Command, []string) error {
		return handleWellKnownErrors(action(b.cmd.Context(), b.cmd.OutOrStdout()))
	}
	return &b.cmd
}

func handleWellKnownErrors(err error) error {
	if err == nil {
		return nil
	}

	if errors.IsUnauthorized(err) {
		fmt.Println("Please check your kubeconfig:", err)
		return nil
	}

	return err
}
