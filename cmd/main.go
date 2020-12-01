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

package main

import (
	"context"
	"errors"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/DevopsArtFactory/klocust/cmd/app"
)

func main() {
	if err := app.Run(os.Stdout, os.Stderr); err != nil {
		if errors.Is(err, context.Canceled) {
			logrus.Debugf("ignore error since context is cancelled: %s", err)
		} else {
			logrus.Errorf(err.Error())
			os.Exit(1)
		}
	}
}
