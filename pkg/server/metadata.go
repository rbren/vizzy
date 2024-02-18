package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rbren/go-prompter/pkg/files"
	"github.com/sirupsen/logrus"

	"github.com/rbren/vizzy/pkg/keys"
)

func getMetadata(c *gin.Context) {
	projectID := c.Param("projectID")
	if projectID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Project ID is required"})
		return
	}
	s3 := files.GetFileManager()
	var metadata map[string]interface{}
	err := s3.ReadJSON(keys.GetMetadataKey(projectID), &metadata)
	if err != nil {
		logrus.WithError(err).Errorf("error getting metatdata from s3 for project %s", projectID)
		c.JSON(http.StatusNotFound, gin.H{"error": "project metadata not found"})
		return
	}
	c.JSON(http.StatusOK, metadata)
}
