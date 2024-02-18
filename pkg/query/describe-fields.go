package query

import (
	"encoding/json"

	"github.com/rbren/go-prompter/pkg/prompt"
)

// DescribeData fills out the template with the provided data and sends the prompt.
func (c *Engine) DescribeFields(desc DataDescription, sample string) (string, error) {
	if len(sample) > maxSampleLength {
		sample = sample[:maxSampleLength]
	}
	b, err := json.Marshal(desc)
	if err != nil {
		return "", err
	}
	var descriptionMap map[string]interface{}
	err = json.Unmarshal(b, &descriptionMap)
	if err != nil {
		return "", err
	}
	descriptionMap["sample"] = sample
	data := map[string]interface{}{
		"data": descriptionMap,
	}

	resp, err := c.QueryWithTemplate("describe-fields", data)
	if err != nil {
		return "", err
	}
	return prompt.ExtractCode(resp)
}
