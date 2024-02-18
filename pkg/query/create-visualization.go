package query

import (
	"encoding/json"
	"os"

	"github.com/rbren/go-prompter/pkg/prompt"
)

type Visualization struct {
	Title string `json:"title"`
	Code  string `json:"code"`
}

const maxSampleLength = 1000

// CreateVisualization creates a visualization from a data description
func (c *Engine) CreateVisualization(userRequest string, desc DataDescription, fieldsMetadata map[string]interface{}, sample string, prev *Visualization) (*Visualization, error) {
	if len(sample) > maxSampleLength {
		sample = sample[:maxSampleLength]
	}
	b, err := json.Marshal(desc)
	if err != nil {
		return nil, err
	}
	var descriptionMap map[string]interface{}
	err = json.Unmarshal(b, &descriptionMap)
	if err != nil {
		return nil, err
	}
	descriptionMap["fieldsMetadata"] = fieldsMetadata
	descriptionMap["sample"] = sample

	data := map[string]interface{}{
		"prompt": userRequest,
		"data":   descriptionMap,
	}
	doPlan := os.Getenv("PLAN_IN_SEPARATE_QUERY") == "true"
	if doPlan && prev == nil {
		plan, err := c.QueryWithTemplate("plan-visualization", data)
		if err != nil {
			return nil, err
		}
		data["plan"] = plan
	}

	template := "create-visualization"
	if prev != nil {
		b, err := json.Marshal(prev)
		if err != nil {
			return nil, err
		}
		prevMap := map[string]interface{}{}
		err = json.Unmarshal(b, &prevMap)
		if err != nil {
			return nil, err
		}
		data["current"] = prevMap
		template = "update-visualization"
	}

	resp, err := c.QueryWithTemplate(template, data)
	if err != nil {
		return nil, err
	}
	code, err := prompt.ExtractCode(resp)
	if err != nil {
		return nil, err
	}
	title, err := prompt.ExtractTitle(resp)
	if err != nil {
		return nil, err
	}
	if title == "" {
		if prev != nil {
			title = prev.Title
		} else {
			title = "Untitled"
		}
	}
	return &Visualization{
		Code: code, Title: title,
	}, nil
}
