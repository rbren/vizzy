package server

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/rbren/go-prompter/pkg/files"

	"github.com/rbren/vizzy/pkg/keys"
	"github.com/rbren/vizzy/pkg/query"
)

func initialDataAnalsys(projectID string, queryEngine *query.Engine) (*query.DataDescription, error) {
	s3 := files.GetFileManager()
	rawData, err := s3.ReadFile(keys.GetDataKey(projectID))
	if err != nil {
		logrus.Error(projectID, err)
		return nil, err
	}
	response, err := queryEngine.DescribeData(string(rawData))
	if err != nil {
		logrus.Error(projectID, err)
		return nil, err
	}
	b, err := json.Marshal(response)
	if err != nil {
		logrus.Error(projectID, err)
		return nil, err
	}
	return &response, s3.WriteFile(keys.GetMetadataKey(projectID), b)
}

func analyzeData(c *gin.Context) {
	projectID := c.Param("projectID")
	openai, err := getOpenAIClient(c)
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	desc, err := initialDataAnalsys(projectID, openai)
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, desc)
}
