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
	"bytes"
	"embed"
	"maps"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
	"gopkg.in/yaml.v2"

	"github.com/DevopsArtFactory/klocust/pkg/schemas"
	"github.com/DevopsArtFactory/klocust/pkg/util"
)

//go:embed _default_templates
var defaultTemplates embed.FS

func renderValuesFile(valuesTemplatePath string, valuesFilePath string, value schemas.LocustValues) (string, error) {
	if util.IsFileExists(valuesFilePath) {
		return "", NewFileExistsError(valuesFilePath)
	}

	t := template.Must(
		template.New(filepath.Base(valuesTemplatePath)).Funcs(sprig.TxtFuncMap()).ParseFS(defaultTemplates, valuesTemplatePath))

	f, err := os.Create(valuesFilePath)
	if err != nil {
		return "", err
	}

	if err := t.Execute(f, value); err != nil {
		return "", err
	}

	return valuesFilePath, err
}

func toYAML(v any) string {
	data, err := yaml.Marshal(v)
	if err != nil {
		// Swallow errors inside of a template.
		return ""
	}
	return strings.TrimSuffix(string(data), "\n")
}

func readFromFile(filename string) string {
	b, err := os.ReadFile(filename)
	if err != nil {
		return ""
	}
	return string(b)
}

func getFileSha256Checksum(filename string) string {
	checksum, err := util.GetSha256Checksum(filename)
	if err != nil {
		return ""
	}
	return checksum
}

func customFuncMap() template.FuncMap {
	f := sprig.TxtFuncMap()
	extra := template.FuncMap{
		"toYaml":                toYAML,
		"readFromFile":          readFromFile,
		"getLocustFilename":     getLocustFilename,
		"getFileSha256Checksum": getFileSha256Checksum,
	}
	maps.Copy(f, extra)
	return f
}

func renderTemplateToBuf(tmplFilepath string, values schemas.LocustValues) (*bytes.Buffer, error) {
	filename := filepath.Base(tmplFilepath)

	t := template.Must(
		template.New(filename).Funcs(customFuncMap()).ParseFS(defaultTemplates, tmplFilepath))

	var buf bytes.Buffer
	if err := t.Execute(&buf, values); err != nil {
		return nil, err
	}

	if buf.Len() == 0 {
		return nil, nil
	}

	return &buf, nil
}
