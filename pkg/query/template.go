package query

import (
	"embed"
	"fmt"
	"strings"
	"text/template"
)

//go:embed prompts/*
var promptsFS embed.FS

func fillTemplate(fileName string, data map[string]interface{}) (string, error) {
	// Parse the template
	tmpl, err := template.New("tmpl").Funcs(template.FuncMap{
		"formatValue": formatValue,
	}).ParseFS(promptsFS, "prompts/*.md")
	if err != nil {
		return "", err
	}
	tmpl = tmpl.Lookup(fileName + ".md")
	if tmpl == nil {
		return "", fmt.Errorf("template %s not found", fileName)
	}

	// Fill out the template with the provided data
	var filledTemplate strings.Builder
	err = tmpl.Execute(&filledTemplate, data)
	if err != nil {
		return "", err
	}
	return filledTemplate.String(), nil
}

func formatValue(value interface{}) string {
	switch value := value.(type) {
	case string:
		if len(value) > 50 {
			return fmt.Sprintf("%s...", value[:50])
		}
		return value
	case []interface{}:
		strs := make([]string, len(value))
		for i, v := range value {
			strs[i] = formatValue(v)
		}
		return strings.Join(strs, ", ")
	default:
		return fmt.Sprintf("%v", value)
	}
}
