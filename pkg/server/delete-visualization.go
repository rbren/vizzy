package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rbren/go-prompter/pkg/files"
	"github.com/sirupsen/logrus"
)

func deleteVisualization(c *gin.Context) {
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

	s3 := files.GetFileManager()
	vizKey := "projects/" + projectID + "/visualizations/" + visualizationID + "/"
	err := s3.DeleteRecursive(vizKey)
	if err != nil {
		logrus.WithError(err).Error("Failed to delete visualization")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}
