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
	"testing"

	"github.com/DevopsArtFactory/klocust/pkg/schemas"
)

func Test_SetDefault(t *testing.T) {
	var values schemas.LocustValues

	SetDefaultToValues(&values)
	if values.Worker.Image != DefaultDockerImage {
		t.Errorf("default value should be applied for worker node image")
	}

	if values.Main.Image != DefaultDockerImage {
		t.Errorf("default value should be applied for main node image")
	}

	testImage := "test:latest"
	values.Worker.Image = testImage
	values.Main.Image = testImage
	SetDefaultToValues(&values)
	if values.Worker.Image != testImage {
		t.Errorf("default value should not be applied for worker node image")
	}

	if values.Main.Image != testImage {
		t.Errorf("default value should not be applied for main node image")
	}
}
