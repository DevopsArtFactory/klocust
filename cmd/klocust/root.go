package klocust

import (
	"context"
	"io"
	"os"

	"github.com/spf13/cobra"
)

func doRoot(_ context.Context, _ io.Writer, cmd *cobra.Command, _ []string) error {
	return cmd.Help()
}

var rootCmd = NewCmd("klocust").
	WithDescription("klocust - A command-line tool for managing Locust distributed load testing on Kubernetes").
	ExactArgs(0, doRoot)

func init() {
	rootCmd.AddCommand(NewCmdList())
	rootCmd.AddCommand(NewInitCmd())
	rootCmd.AddCommand(NewApplyCMD())
	rootCmd.AddCommand(NewCmdCompletion())
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
