// completion from https://github.com/GoogleContainerTools/skaffold
package klocust

import (
	"context"
	"io"
	"os"

	"github.com/spf13/cobra"
)

const (
	longDescription = `Outputs shell completion for the given shell (bash or zsh)

This depends on the bash-completion binary.  Example installation instructions:
OS X:
	$ brew install bash-completion
	$ source $(brew --prefix)/etc/bash_completion
	$ klocust completion bash > ~/.klocust-completion  # for bash users
	$ klocust completion zsh > ~/.klocust-completion   # for zsh users
	$ source ~/.klocust-completion
Ubuntu:
	$ apt-get install bash-completion
	$ source /etc/bash-completion
	$ source <(klocust completion bash) # for bash users
	$ source <(klocust completion zsh)  # for zsh users

Additionally, you may want to output the completion to a file and source in your .bashrc
`
	zshCompdef = "\ncompdef _klocust klocust\n"
)

func completion(_ context.Context, _ io.Writer, cmd *cobra.Command, args []string) error {
	switch args[0] {
	case "bash":
		return findRootCmd(cmd).GenBashCompletion(os.Stdout)
	case "zsh":
		return runCompletionZsh(cmd, os.Stdout)
	}
	return nil
}

// NewCmdCompletion returns the cobra command that outputs shell completion code
func NewCmdCompletion() *cobra.Command {
	return NewCmd("completion").
		WithDescription("Output shell completion for the given shell (bash or zsh)").
		ExactArgs(1, completion)
}

func runCompletionZsh(cmd *cobra.Command, out io.Writer) error {
	if err := findRootCmd(cmd).GenZshCompletion(out); err != nil {
		return err
	}
	if _, err := io.WriteString(out, zshCompdef); err != nil {
		return err
	}
	return nil
}

func findRootCmd(cmd *cobra.Command) *cobra.Command {
	parent := cmd
	for parent.HasParent() {
		parent = parent.Parent()
	}
	return parent
}
