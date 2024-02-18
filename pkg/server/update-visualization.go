package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/rbren/go-prompter/pkg/files"

	"github.com/rbren/vizzy/pkg/keys"
	"github.com/rbren/vizzy/pkg/query"
)

type UpdateVisualizationRequestBody struct {
	Code string `json:"code"`
}

func updateVisualization(c *gin.Context) {
	s3 := files.GetFileManager()

	projectID := c.Param("projectID")
	if projectID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Project ID is required"})
		return
	}
	visualizationID := c.Param("visualizationID")
	if visualizationID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Visualization ID is required"})
		return
	}
	versionIDStr := c.Param("versionID")
	if versionIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Version ID is required"})
		return
	}
	versionID, err := strconv.ParseInt(versionIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Version ID must be an integer"})
		return
	}
	nextVersionID := versionID + 1

	baseVisualization := map[string]interface{}{}
	err = s3.ReadJSON(keys.GetVisualizationVersionKey(projectID, visualizationID, int(versionID)), &baseVisualization)
	if err != nil {
		logrus.WithError(err).Errorf("error getting current version from s3 for project %s", projectID)
		c.JSON(http.StatusNotFound, gin.H{"error": "project visualization not found"})
		return
	}
	prev := query.Visualization{
		Title: baseVisualization["title"].(string),
		Code:  baseVisualization["visualization"].(string),
	}

	newTitle := c.Query("title")
	if newTitle == "" && baseVisualization["title"] != nil {
		newTitle = baseVisualization["title"].(string)
	}
	newCode := ""
	prompt := ""
	if c.Request.ContentLength > 0 {
		reqBody := UpdateVisualizationRequestBody{}
		err := c.BindJSON(&reqBody)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}
		newCode = reqBody.Code
		prompt = "[manual edit]"
	}

	if newCode == "" {
		prompt = c.Query("prompt")
		if prompt == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Prompt is required"})
			return
		}

		oaiClient, err := getOpenAIClient(c)
		if err != nil {
			logrus.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var metadata query.DataDescription
		err = s3.ReadJSON(keys.GetMetadataKey(projectID), &metadata)
		if err != nil {
			logrus.WithError(err).Errorf("error getting metatdata from s3 for project %s", projectID)
			c.JSON(http.StatusNotFound, gin.H{"error": "project metadata not found"})
			return
		}

		var fieldsCode map[string]interface{}
		err = s3.ReadJSON(keys.GetFieldsCodeKey(projectID), &fieldsCode)
		if err != nil {
			logrus.WithError(err).Errorf("error getting fields metatdata from s3 for project %s", projectID)
			c.JSON(http.StatusNotFound, gin.H{"error": "project fields metadata not found"})
			return
		}
		fieldsMetadata, _ := fieldsCode["metadata"].(map[string]interface{})

		data, err := s3.ReadFile(keys.GetDataKey(projectID))
		if err != nil {
			logrus.WithError(err).Errorf("error getting data from s3 for project %s", projectID)
			c.JSON(http.StatusNotFound, gin.H{"error": "project data not found"})
			return
		}

		viz, err := oaiClient.CreateVisualization(prompt, metadata, fieldsMetadata, string(data), &prev)
		if err != nil {
			logrus.WithError(err).Errorf("error creating visualization for project %s", projectID)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error creating visualization"})
			return
		}
		newCode = viz.Code
		newTitle = viz.Title
	}
	data := map[string]string{
		"id":            visualizationID,
		"prompt":        prompt,
		"visualization": newCode,
		"title":         newTitle,
	}

	err = s3.WriteJSON(keys.GetVisualizationVersionKey(projectID, visualizationID, int(nextVersionID)), data)
	if err != nil {
		logrus.WithError(err).Errorf("error writing visualization to s3 for project %s", projectID)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error creating visualization"})
		return
	}
	c.JSON(http.StatusOK, data)
}
