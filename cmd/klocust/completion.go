package klocust

import (
	"fmt"
	"io"
	"log"
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

func completion(cmd *cobra.Command, args []string) {
	switch args[0] {
	case "bash":
		if err := findRootCmd(cmd).GenBashCompletion(os.Stdout); err != nil {
			log.Fatal(err)
		}
	case "zsh":
		if err := runCompletionZsh(cmd, os.Stdout); err != nil {
			log.Fatal(err)
		}
	}
}

// NewCmdCompletion returns the cobra command that outputs shell completion code
func NewCmdCompletion() *cobra.Command {
	return &cobra.Command{
		Use: "completion (bash|zsh)",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("requires 1 arg, found %d", len(args))
			}
			return cobra.OnlyValidArgs(cmd, args)
		},
		ValidArgs: []string{"bash", "zsh"},
		Short:     "Output shell completion for the given shell (bash or zsh)",
		Long:      longDescription,
		Run:       completion,
	}
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
