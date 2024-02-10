package llm

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/rbren/vizzy/pkg/files"
)

// OpenAIRequest represents the request body for OpenAI API.
type OpenAIRequest struct {
	Model    string          `json:"model"`
	Messages []OpenAIMessage `json:"messages"`
	Seed     int             `json:"seed,omitempty"`
}

// OpenAIMessage represents a message in the request body for OpenAI API.
type OpenAIMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// OpenAIResponse represents the response from OpenAI API.
type OpenAIResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
	Error map[string]interface{} `json:"error"`
}

// OpenAIClient holds the information needed to make requests to the OpenAI API.
type OpenAIClient struct {
	APIKey      string
	Model       string
	FileManager files.FileManager
	Seed        int
}

// NewOpenAIClient creates a new OpenAI API client.
func NewOpenAIClient(apiKey, model string) *OpenAIClient {
	return &OpenAIClient{
		APIKey: apiKey,
		Model:  model,
	}
}

func (c *OpenAIClient) Copy() Client {
	return &OpenAIClient{
		APIKey:      c.APIKey,
		Model:       c.Model,
		FileManager: c.FileManager,
		Seed:        c.Seed,
	}
}

func (c *OpenAIClient) SetDebugFileManager(mgr files.FileManager) {
	c.FileManager = mgr
}

// Query sends a prompt to the OpenAI API and returns the response.
func (c *OpenAIClient) Query(id string, prompt string) (string, error) {
	systemPrompt := "The following is a conversation with an AI assistant."

	if c.FileManager != nil {
		err := c.FileManager.WriteFile(id+"/request.md", []byte(prompt))
		if err != nil {
			return "", err
		}
	}

	requestBody, err := json.Marshal(OpenAIRequest{
		Seed: c.Seed,
		Messages: []OpenAIMessage{{
			Role:    "system",
			Content: systemPrompt,
		}, {
			Role:    "user",
			Content: prompt,
		}},
		Model: c.Model,
	})
	if err != nil {
		return "", err
	}
	logrus.Debugf("request: %s", prompt)

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(requestBody))
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

	var response OpenAIResponse
	logrus.Debugf("response: %s", string(body))
	if err := json.Unmarshal(body, &response); err != nil {
		return "", err
	}

	if response.Error != nil {
		logrus.Errorf("error from OpenAI: %v", response.Error)
		message, ok := response.Error["message"].(string)
		if ok {
			return "", errors.New(message)
		} else {
			return "", errors.New("unknown error from OpenAI")
		}
	}

	if len(response.Choices) == 0 {
		return "", fmt.Errorf("no response from OpenAI")
	}

	if c.FileManager != nil {
		err := c.FileManager.WriteFile(id+"/response.md", []byte(response.Choices[0].Message.Content))
		if err != nil {
			return "", err
		}
	}
	return response.Choices[0].Message.Content, nil
}
