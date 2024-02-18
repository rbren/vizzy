package server

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rbren/go-prompter/pkg/llm"
	"github.com/rbren/go-prompter/pkg/prompt"

	"github.com/rbren/vizzy/pkg/query"
)

func getOpenAIClient(c *gin.Context) (*query.Engine, error) {
	apiKey := c.GetHeader("X-OPENAI-API-KEY")
	if apiKey == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "X-OPENAI-API-KEY is required"})
		return nil, errors.New("OpenAI API key is required")
	}

	modelName := c.GetHeader("X-OPENAI-MODEL")
	if modelName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "X-OPENAI-MODEL header is required"})
		return nil, errors.New("OpenAI model name is required")
	}

	llmClient := llm.NewOpenAIClient(apiKey, modelName)
	return &query.Engine{
		Prompter: &prompt.Engine{
			LLM: llmClient,
		},
	}, nil
}
