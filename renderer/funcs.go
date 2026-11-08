package renderer

import (
	"fmt"
	"html/template"
	"os"
)

var (
	funcs = template.FuncMap{
		"includeHTML": func(filePath string) (template.HTML, error) {
			bytes, err := os.ReadFile(filePath)
			if err != nil {
				return "", fmt.Errorf("failed to read file %s: %w", filePath, err)
			}

			return template.HTML(bytes), nil
		},
	}
)
