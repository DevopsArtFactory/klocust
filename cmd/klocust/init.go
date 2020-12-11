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
	"context"
	"io"

	"github.com/spf13/cobra"

	"github.com/DevopsArtFactory/klocust/internal/klocust"
)

func doInit(_ context.Context, out io.Writer, _ *cobra.Command, args []string) error {
	locustName := args[0]
	return klocust.InitLocust(out, opts.Namespace, locustName)
}

func NewInitCmd() *cobra.Command {
	return NewCmd("init").
		WithDescription("initiate klocust cluster").
		WithCommonFlags().
		ExactArgs(1, doInit)
}
