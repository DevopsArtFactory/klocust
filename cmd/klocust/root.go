/*
Copyright 2020 The klocust Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package klocust

import (
	"fmt"
	"io"
	"os"

	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/color"
	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"k8s.io/kubectl/pkg/util/templates"

	"github.com/DevopsArtFactory/klocust/pkg/klocust"
	"github.com/DevopsArtFactory/klocust/pkg/version"
)

var (
	cfgFile string
	v       string

	// User can set default printer
	defaultColor int
	forceColors  bool
)

// Get root command
func NewRootCommand(out, stderr io.Writer) *cobra.Command {
	cobra.OnInitialize(initConfig)
	rootCmd := &cobra.Command{
		Use:           "klocust",
		Short:         "klocust - A command-line tool for managing Locust distributed load testing on Kubernetes",
		Long:          "klocust - A command-line tool for managing Locust distributed load testing on Kubernetes",
		SilenceErrors: true,
		SilenceUsage:  true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Setup logs
			if err := SetUpLogs(stderr, v); err != nil {
				return err
			}

			out = color.SetupColors(out, defaultColor, forceColors)
			cmd.Root().SetOutput(out)

			version := version.Get()

			logrus.Infof("klocust %+v", version)

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	// Group by commands
	groups := templates.CommandGroups{
		{
			Message: "Initiates klocust cluster",
			Commands: []*cobra.Command{
				NewInitCmd(),
			},
		},
		{
			Message: "Manage klocust clusters",
			Commands: []*cobra.Command{
				NewListCmd(),
				NewApplyCmd(),
				NewDeleteCmd(),
			},
		},
		{
			Message: "Helper operations of klocust",
			Commands: []*cobra.Command{
				NewCompletionCmd(),
			},
		},
	}

	groups.Add(rootCmd)
	rootCmd.AddCommand(NewVersionCmd())

	rootCmd.PersistentFlags().StringVarP(&v, "verbosity", "v", klocust.DefaultLogLevel.String(), "Log level (debug, info, warn, error, fatal, panic)")

	templates.ActsAsRootCommand(rootCmd, nil, groups...)

	return rootCmd
}

// SetUpLogs set logrus log format
func SetUpLogs(stdErr io.Writer, level string) error {
	logrus.SetOutput(stdErr)
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		return fmt.Errorf("parsing log level: %w", err)
	}
	logrus.SetLevel(lvl)
	return nil
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
	}

	viper.AutomaticEnv() // read in environment variables that match
}
