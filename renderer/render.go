package renderer

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

type (
	templateData map[string]any
)

func Render(dataPath, outPath string, tmplFiles ...string) error {
	dataBytes, err := os.ReadFile(dataPath)
	if err != nil {
		return fmt.Errorf("failed to read data file %s: %w", dataPath, err)
	}

	var data templateData
	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		return fmt.Errorf("failed to parse data file %s: %w", dataPath, err)
	}

	outFile, err := os.Create(outPath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", outPath, err)
	}
	defer func() {
		_ = outFile.Close()
	}()

	if len(tmplFiles) == 0 {
		return nil
	}

	tmpl := template.Must(template.New(filepath.Base(tmplFiles[0])).Funcs(template.FuncMap{
		"includeHTML": func(filePath string) (template.HTML, error) {
			bytes, err := os.ReadFile(filePath)
			if err != nil {
				return "", fmt.Errorf("failed to read file %s: %w", filePath, err)
			}

			return template.HTML(bytes), nil
		},
	}).ParseFiles(tmplFiles...))

	return tmpl.Execute(outFile, data)
}
