package renderer

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"
)

type (
	templateData map[string]any
)

func Render(tmplPath, dataPath, outPath string) error {
	dataBytes, err := os.ReadFile(dataPath)
	if err != nil {
		return fmt.Errorf("failed to read data file %s: %w", dataPath, err)
	}

	var data templateData
	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		return fmt.Errorf("failed to parse data file %s: %w", dataPath, err)
	}

	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return fmt.Errorf("failed to parse template file %s: %w", tmplPath, err)
	}

	outFile, err := os.Create(outPath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", outPath, err)
	}
	defer func() {
		_ = outFile.Close()
	}()

	return tmpl.Execute(outFile, data)
}
