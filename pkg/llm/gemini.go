package llm

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/rbren/vizzy/pkg/files"
)

const URL = "https://generativelanguage.googleapis.com/v1beta/models/gemini-pro:generateContent"

type GeminiRequest struct {
	Contents []struct {
		Parts []struct {
			Text string `json:"text"`
		} `json:"parts"`
	} `json:"contents"`
}

type GeminiResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	}
}

type Gemini struct {
	APIKey      string
	FileManager files.FileManager
}

func NewGeminiClient(apiKey string) *Gemini {
	return &Gemini{
		APIKey: apiKey,
	}
}

func (g *Gemini) Copy() Client {
	return &Gemini{
		APIKey:      g.APIKey,
		FileManager: g.FileManager,
	}
}

func (g *Gemini) SetDebugFileManager(fm files.FileManager) {
	g.FileManager = fm
}

func (g *Gemini) Query(id, prompt string) (string, error) {
	if g.FileManager != nil {
		err := g.FileManager.WriteFile(id+"/request.md", []byte(prompt))
		if err != nil {
			return "", err
		}
	}
	request := GeminiRequest{
		Contents: []struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		}{
			{Parts: []struct {
				Text string `json:"text"`
			}{{prompt}}},
		},
	}

	reqBody, err := json.Marshal(request)
	if err != nil {
		return "", err
	}

	url := URL + "?key=" + g.APIKey
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

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
	if resp.StatusCode != 200 {
		logrus.Errorf("Gemini Response: %s", resp.Status)
		return "", errors.New("non-200 status code")
	}

	var response GeminiResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}
	if len(response.Candidates) == 0 {
		return "", errors.New("no candidates")
	}
	if len(response.Candidates[0].Content.Parts) == 0 {
		return "", errors.New("no parts")
	}
	text := response.Candidates[0].Content.Parts[0].Text
	if g.FileManager != nil {
		err := g.FileManager.WriteFile(id+"/response.md", []byte(text))
		if err != nil {
			return "", err
		}
	}
	return text, nil
}
