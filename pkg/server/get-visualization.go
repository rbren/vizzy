package server

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/rbren/vizzy/pkg/files"
)

func getVisualization(c *gin.Context) {
	versionStr := c.Param("version")
	version, err := strconv.Atoi(versionStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "version must be an integer"})
		return
	}
	visualizationID := c.Param("visualizationID")
	projectID := c.Param("projectID")
	if projectID == "" || visualizationID == "" {
		c.JSON(400, gin.H{"error": "you must specify a projectID, visualizationID, and version"})
		return
	}
	s3 := files.GetFileManager()
	viz := map[string]interface{}{}
	err = s3.ReadJSON(files.GetVisualizationVersionKey(projectID, visualizationID, version), &viz)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, viz)
}
