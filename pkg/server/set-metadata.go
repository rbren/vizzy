package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/rbren/go-prompter/pkg/files"

	"github.com/rbren/vizzy/pkg/keys"
	"github.com/rbren/vizzy/pkg/query"
)

func setMetadata(c *gin.Context) {
	projectID := c.Param("projectID")
	if projectID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Project ID is required"})
		return
	}

	bodyData := query.DataDescription{}
	if err := c.BindJSON(&bodyData); err != nil {
		logrus.WithError(err).Errorf("error binding metadata json for project %s", projectID)
		c.JSON(http.StatusBadRequest, gin.H{"error": "error parsing metadata request body"})
		return
	}

	s3 := files.GetFileManager()
	err := s3.WriteJSON(keys.GetMetadataKey(projectID), bodyData)
	if err != nil {
		logrus.WithError(err).Errorf("error getting metatdata from s3 for project %s", projectID)
		c.JSON(http.StatusNotFound, gin.H{"error": "project metadata not found"})
		return
	}
	c.JSON(http.StatusOK, bodyData)
}
