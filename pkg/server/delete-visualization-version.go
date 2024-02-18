package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/rbren/go-prompter/pkg/files"

	"github.com/rbren/vizzy/pkg/keys"
)

func deleteVisualizationVersion(c *gin.Context) {
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
	logrus.Infof("deleting visualization version %d for project %s", versionID, projectID)
	err = s3.DeleteFile(keys.GetVisualizationVersionKey(projectID, visualizationID, int(versionID)))
	if err != nil {
		logrus.WithError(err).Errorf("error deleting visualization version from s3 for project %s", projectID)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error deleting visualization version"})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"success": true})
}
