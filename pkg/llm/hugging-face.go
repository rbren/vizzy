package llm

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/rbren/vizzy/pkg/files"
)

// HuggingFaceRequest represents the request body for HuggingFace API.
type HuggingFaceRequest struct {
	Inputs string `json:"inputs"`
}

type HuggingFaceResponse struct {
	Error string `json:"error"`
}

// HuggingFaceClient holds the information needed to make requests to the HuggingFace API.
type HuggingFaceClient struct {
	URL         string
	APIKey      string
	FileManager files.FileManager
}

// NewHuggingFaceClient creates a new HuggingFace API client.
func NewHuggingFaceClient(apiKey, url string) *HuggingFaceClient {
	return &HuggingFaceClient{
		APIKey: apiKey,
		URL:    url,
	}
}

func (c *HuggingFaceClient) Copy() Client {
	return &HuggingFaceClient{
		APIKey:      c.APIKey,
		URL:         c.URL,
		FileManager: c.FileManager,
	}
}

func (c *HuggingFaceClient) SetDebugFileManager(mgr files.FileManager) {
	c.FileManager = mgr
}

// Query sends a prompt to the HuggingFace API and returns the response.
func (c *HuggingFaceClient) Query(id string, prompt string) (string, error) {
	systemPrompt := "<s>Source: system\n\nThe following is a conversation with an AI assistant."
	stepPrompt := " <step> Source: user\n\n" + prompt
	destPrompt := " <step> Source: assistant\nDestination: user\n\n"
	finalPrompt := systemPrompt + stepPrompt + destPrompt

	if c.FileManager != nil {
		err := c.FileManager.WriteFile(id+"/request.md", []byte(prompt))
		if err != nil {
			return "", err
		}
	}

	requestBody, err := json.Marshal(HuggingFaceRequest{
		Inputs: finalPrompt,
	})
	if err != nil {
		return "", err
	}
	logrus.Debugf("request: %s", prompt)

	req, err := http.NewRequest("POST", c.URL, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.APIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	logrus.Infof("response: %s", string(body))

	var response []map[string]interface{}
	logrus.Debugf("response: %s", string(body))
	if err := json.Unmarshal(body, &response); err != nil {
		return "", err
	}

	out := response[0]["generated_text"].(string)
	out = strings.TrimPrefix(out, prompt)

	if c.FileManager != nil {
		err := c.FileManager.WriteFile(id+"/response.md", []byte(out))
		if err != nil {
			return "", err
		}
	}
	return out, nil
}
