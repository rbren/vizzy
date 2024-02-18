package query

import (
	"fmt"
	"strings"
	"embed"
	"text/template"

	"github.com/rbren/go-prompter/pkg/prompt"
)

//go:embed prompts/*
var promptsFS embed.FS

func init() {
	prompt.SetFS(&promptsFS)
	prompt.SetTemplateFuncMap(template.FuncMap{
		"formatValue": formatValue,
	})
}

type Engine struct {
	Prompter *prompt.Engine
}

func New() *Engine {
	return &Engine{
		Prompter: prompt.New(),
	}
}

func (c *Engine) WithSession(id string) (*Engine) {
	return &Engine{
		Prompter: c.Prompter.WithSession(id),
	}
}

func (c *Engine) QueryWithTemplate(template string, data map[string]interface{}) (string, error) {
	return c.Prompter.PromptWithTemplate(template, data)
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
