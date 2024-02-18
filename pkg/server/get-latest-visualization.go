package server

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rbren/go-prompter/pkg/files"
	"github.com/sirupsen/logrus"

	"github.com/rbren/vizzy/pkg/keys"
)

func getLatestVisualization(c *gin.Context) {
	visualizationID := c.Param("visualizationID")
	projectID := c.Param("projectID")
	if projectID == "" || visualizationID == "" {
		c.JSON(400, gin.H{"error": "you must specify a projectID, visualizationID, and version"})
		return
	}
	s3 := files.GetFileManager()
	vizKey := "projects/" + projectID + "/visualizations/" + visualizationID + "/"
	versions, err := s3.ListFilesRecursive(vizKey)
	if err != nil {
		logrus.WithError(err).Error("error listing versions")
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	logrus.Infof("Found %d versions for %s", len(versions), vizKey)
	largest := 0
	for _, ver := range versions {
		logrus.Info(ver)
		parts := strings.Split(ver, "/")
		if len(parts) == 0 {
			logrus.Warnf("invalid version %s", ver)
			continue
		}
		ver := strings.TrimSuffix(parts[len(parts)-1], ".json")
		verNum, err := strconv.ParseInt(ver, 10, 64)
		if err != nil {
			logrus.WithError(err).Warnf("invalid version parse %s", ver)
			continue
		}
		if int(verNum) > largest {
			largest = int(verNum)
		}
	}
	viz := map[string]interface{}{}
	err = s3.ReadJSON(keys.GetVisualizationVersionKey(projectID, visualizationID, largest), &viz)
	if err != nil {
		logrus.WithError(err).Errorf("error reading visualization")
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	viz["version"] = largest
	c.JSON(200, viz)
}
