package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/rbren/go-prompter/pkg/files"

	"github.com/rbren/vizzy/pkg/keys"
)

func getData(c *gin.Context) {
	projectID := c.Param("projectID")
	if projectID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Project ID is required"})
		return
	}
	s3 := files.GetFileManager()
	data, err := s3.ReadFile(keys.GetDataKey(projectID))
	if err != nil {
		logrus.WithError(err).Errorf("error getting tdata from s3 for project %s", projectID)
		c.JSON(http.StatusNotFound, gin.H{"error": "project data not found"})
		return
	}
	c.String(http.StatusOK, string(data))
}
