package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/rbren/vizzy/pkg/files"
	"github.com/rbren/vizzy/pkg/query"
)

func describeFields(c *gin.Context) {
	projectID := c.Param("projectID")
	if projectID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Project ID is required"})
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

	data, err := s3.ReadFile(files.GetDataKey(projectID))
	if err != nil {
		logrus.WithError(err).Errorf("error getting data from s3 for project %s", projectID)
		c.JSON(http.StatusNotFound, gin.H{"error": "project data not found"})
		return
	}

	code, err := oaiClient.DescribeFields(metadata, string(data))
	if err != nil {
		logrus.WithError(err).Errorf("error describing fields for project %s", projectID)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error describing fields"})
		return
	}

	resp := map[string]interface{}{
		"code": code,
	}
	err = s3.WriteJSON(files.GetFieldsCodeKey(projectID), resp)
	if err != nil {
		logrus.WithError(err).Errorf("error writing fields code to s3 for project %s", projectID)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error generating fields code"})
		return
	}
	c.JSON(http.StatusOK, resp)
}
