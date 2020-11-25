package klocust

import (
	"gopkg.in/yaml.v2"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/DevopsArtFactory/klocust/internal/schemas"
	"github.com/DevopsArtFactory/klocust/internal/util"
	"github.com/Masterminds/sprig"
)

func renderValuesFile(valuesTemplatePath string, valuesFilePath string, value schemas.LocustValues) (string, error) {
	if util.IsFileExists(valuesFilePath) {
		return "", NewFileExistsError(valuesFilePath)
	}

	t := template.Must(
		template.New("values.yaml").Funcs(sprig.FuncMap()).ParseFiles(valuesTemplatePath))

	f, err := os.Create(valuesFilePath)
	if err != nil {
		return "", err
	}

	if err := t.Execute(f, value); err != nil {
		return "", err
	}

	return valuesFilePath, nil
}

func toYAML(v interface{}) string {
	data, err := yaml.Marshal(v)
	if err != nil {
		// Swallow errors inside of a template.
		return ""
	}
	return strings.TrimSuffix(string(data), "\n")
}

func customFuncMap() template.FuncMap {
	f := sprig.FuncMap()
	extra := template.FuncMap{
		"toYaml": toYAML,
	}
	for k, v := range extra {
		f[k] = v
	}
	return f
}

func renderTemplateFile(tmplFilepath string, projectFilepath string, values schemas.LocustValues) (string, error) {
	filename := filepath.Base(tmplFilepath)

	t := template.Must(
		template.New(filename).Funcs(customFuncMap()).ParseFiles(tmplFilepath))

	f, err := os.Create(projectFilepath)
	if err != nil {
		return "", err
	}

	if err := t.Execute(f, values); err != nil {
		return "", err
	}

	return projectFilepath, nil
}
