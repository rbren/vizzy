package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rbren/go-prompter/pkg/files"
	"github.com/sirupsen/logrus"
)

func listVisualizations(c *gin.Context) {
	projectID := c.Param("projectID")
	if projectID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Project ID is required"})
		return
	}
	s3 := files.GetFileManager()
	projectKey := "projects/" + projectID + "/visualizations/"
	files, err := s3.ListDirectories(projectKey)
	if err != nil {
		logrus.WithError(err).Warning("Failed to list files, assuming empty")
		files = []string{}
	}

	c.JSON(http.StatusOK, gin.H{"ids": files})
}
