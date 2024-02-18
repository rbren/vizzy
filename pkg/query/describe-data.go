package query

import (
	"encoding/json"

	"github.com/rbren/go-prompter/pkg/prompt"
)

const charsToAnalyze = 8000

type DataDescription struct {
	Type                    string   `json:"type"`
	Title                   string   `json:"title"`
	Description             string   `json:"description"`
	DataFormat              string   `json:"data_format"`
	Fields                  []string `json:"fields"`
	SuggestedVisualizations []string `json:"suggested_visualizations"`
}

// DescribeData fills out the template with the provided data and sends the prompt.
func (c *Engine) DescribeData(data string) (DataDescription, error) {
	data = data[:min(charsToAnalyze, len(data))]

	desc := DataDescription{}

	resp, err := c.QueryWithTemplate("describe-data", map[string]interface{}{"data": data, "characters": charsToAnalyze})
	if err != nil {
		return desc, err
	}

	err = parseJSONObject(resp, &desc)
	return desc, err
}

func parseJSONObject(body string, v interface{}) error {
	jsonObject, err := prompt.ExtractJSONObject(body)
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(jsonObject), v)
}
