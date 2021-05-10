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

package version

import (
	"fmt"
	"runtime"
)

var version, gitCommit, gitTreeState, buildDate string
var platform = fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)

type Controller struct{}

type Info struct {
	Version      string
	BuildDate    string
	GitCommit    string
	GitTreeState string
	Platform     string
}

func Get() Info {
	return Info{
		Version:      version,
		BuildDate:    buildDate,
		GitCommit:    gitCommit,
		GitTreeState: gitTreeState,
		Platform:     platform,
	}
}

func (v Controller) Print(info Info) error {
	_, err := fmt.Println(info.Version)
	return err
}
