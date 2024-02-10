package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/rbren/vizzy/pkg/files"
	"github.com/rbren/vizzy/pkg/query"
)

func createVisualization(c *gin.Context) {
	projectID := c.Param("projectID")
	if projectID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Project ID is required"})
		return
	}
	prompt := c.Query("prompt")
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

	s3 := files.GetFileManager()
	var metadata query.DataDescription
	err = s3.ReadJSON(files.GetMetadataKey(projectID), &metadata)
	if err != nil {
		logrus.WithError(err).Errorf("error getting metatdata from s3 for project %s", projectID)
		c.JSON(http.StatusNotFound, gin.H{"error": "project metadata not found"})
		return
	}

	var fieldsCode map[string]interface{}
	err = s3.ReadJSON(files.GetFieldsCodeKey(projectID), &fieldsCode)
	if err != nil {
		logrus.WithError(err).Errorf("error getting fields metatdata from s3 for project %s", projectID)
		c.JSON(http.StatusNotFound, gin.H{"error": "project fields metadata not found"})
		return
	}
	fieldsMetadata, _ := fieldsCode["metadata"].(map[string]interface{})

	data, err := s3.ReadFile(files.GetDataKey(projectID))
	if err != nil {
		logrus.WithError(err).Errorf("error getting data from s3 for project %s", projectID)
		c.JSON(http.StatusNotFound, gin.H{"error": "project data not found"})
		return
	}

	vizUUID := files.GenerateUUID()
	viz, err := oaiClient.CreateVisualization(prompt, metadata, fieldsMetadata, string(data), nil)
	if err != nil {
		logrus.WithError(err).Errorf("error creating visualization for project %s", projectID)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error creating visualization"})
		return
	}
	resp := map[string]string{
		"id":            vizUUID,
		"prompt":        prompt,
		"visualization": viz.Code,
		"title":         viz.Title,
	}

	err = s3.WriteJSON(files.GetVisualizationVersionKey(projectID, vizUUID, 1), resp)
	if err != nil {
		logrus.WithError(err).Errorf("error writing visualization to s3 for project %s", projectID)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error creating visualization"})
		return
	}
	c.JSON(http.StatusOK, resp)
}
