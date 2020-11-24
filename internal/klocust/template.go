package klocust

import (
	"github.com/DevopsArtFactory/klocust/internal/schemas"
	"github.com/DevopsArtFactory/klocust/internal/util"
	"github.com/Masterminds/sprig"
	"html/template"
	"os"
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
