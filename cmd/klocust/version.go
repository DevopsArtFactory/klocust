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

	"github.com/DevopsArtFactory/klocust/pkg/version"
)

// Get klocust version
func NewVersionCmd() *cobra.Command {
	return NewCmd("version").
		WithDescription("Print the version information").
		SetAliases([]string{"v"}).
		NoArgs(funcVersion)
}

// funcVersion
func funcVersion(_ context.Context, _ io.Writer) error {
	return version.Controller{}.Print(version.Get())
}
