package query

import (
	"github.com/rbren/vizzy/pkg/llm"
)

type Engine struct {
	LLM llm.Client
}

func (c *Engine) QueryWithTemplate(template string, data map[string]interface{}) (string, error) {
	prompt, err := fillTemplate(template, data)
	if err != nil {
		return "", err
	}
	return c.LLM.Query(template, prompt)
}
