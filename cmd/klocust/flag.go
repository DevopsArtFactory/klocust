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
	"reflect"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type Flag struct {
	Name               string
	Shorthand          string
	Usage              string
	Value              interface{}
	DefValue           interface{}
	DefValuePerCommand map[string]interface{}
	FlagAddMethod      string
	DefinedOn          []string
	Hidden             bool

	pflag *pflag.Flag
}

type CommonFlag struct {
	Namespace string
}

var opts CommonFlag

var flagRegistry = []Flag{
	{
		Name:          "namespace",
		Shorthand:     "n",
		Usage:         "Kubernetes namespace",
		Value:         &opts.Namespace,
		DefValue:      "",
		FlagAddMethod: "StringVar",
		DefinedOn:     []string{"all"},
	},
}

func (fl *Flag) flag() *pflag.Flag {
	if fl.pflag != nil {
		return fl.pflag
	}

	inputs := []interface{}{fl.Value, fl.Name}
	if fl.FlagAddMethod != "Var" {
		inputs = append(inputs, fl.DefValue)
	}
	inputs = append(inputs, fl.Usage)

	fs := pflag.NewFlagSet(fl.Name, pflag.ContinueOnError)

	reflect.ValueOf(fs).MethodByName(fl.FlagAddMethod).Call(reflectValueOf(inputs))
	f := fs.Lookup(fl.Name)
	f.Shorthand = fl.Shorthand
	f.Hidden = fl.Hidden

	fl.pflag = f
	return f
}

func reflectValueOf(values []interface{}) []reflect.Value {
	var results []reflect.Value
	for _, v := range values {
		results = append(results, reflect.ValueOf(v))
	}
	return results
}

// AddFlags adds to the command the common flags that are annotated with the command name.
func AddFlags(cmd *cobra.Command) {
	var flagsForCommand []*Flag

	for i := range flagRegistry {
		fl := &flagRegistry[i]
		if !hasCmdAnnotation(cmd.Use, fl.DefinedOn) {
			continue
		}

		cmd.Flags().AddFlag(fl.flag())

		flagsForCommand = append(flagsForCommand, fl)
	}

	// Apply command-specific default values to flags.
	cmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		// Update default values.
		for _, fl := range flagsForCommand {
			if defValue, present := fl.DefValuePerCommand[cmd.Use]; present {
				if flag := cmd.Flag(fl.Name); !flag.Changed {
					flag.Value.Set(fmt.Sprintf("%v", defValue))
				}
			}
		}

		// Since PersistentPreRunE replaces the parent's PersistentPreRunE,
		// make sure we call it, if it is set.
		if parent := cmd.Parent(); parent != nil {
			if preRun := parent.PersistentPreRunE; preRun != nil {
				if err := preRun(cmd, args); err != nil {
					return err
				}
			} else if preRun := parent.PersistentPreRun; preRun != nil {
				preRun(cmd, args)
			}
		}

		return nil
	}
}

func hasCmdAnnotation(cmdName string, annotations []string) bool {
	for _, a := range annotations {
		if cmdName == a || a == "all" {
			return true
		}
	}
	return false
}
