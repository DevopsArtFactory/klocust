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

package printer

// Use Skaffold printer package
import (
	"io"

	colors "github.com/GoogleContainerTools/skaffold/pkg/skaffold/color"
)

type Printer struct {
	Printer colors.Color
}

var (
	Default     = Printer{Printer: colors.Default}
	LightRed    = Printer{Printer: colors.LightRed}
	LightGreen  = Printer{Printer: colors.LightGreen}
	LightYellow = Printer{Printer: colors.LightYellow}
	LightBlue   = Printer{Printer: colors.LightBlue}
	LightPurple = Printer{Printer: colors.LightPurple}
	Red         = Printer{Printer: colors.Red}
	Green       = Printer{Printer: colors.Green}
	Yellow      = Printer{Printer: colors.Yellow}
	Blue        = Printer{Printer: colors.Blue}
	Purple      = Printer{Printer: colors.Purple}
	Cyan        = Printer{Printer: colors.Cyan}
	White       = Printer{Printer: colors.White}
	None        = Printer{Printer: colors.None}
)

// Fprintln outputs with color output.
func (p Printer) Fprintln(out io.Writer, a ...any) {
	p.Printer.Fprintln(out, a...)
}

// Fprintf outputs with format
func (p Printer) Fprintf(out io.Writer, format string, a ...any) {
	p.Printer.Fprintf(out, format, a...)
}
