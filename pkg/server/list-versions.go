package server

import (
	"sort"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/rbren/vizzy/pkg/files"
)

func listVersions(c *gin.Context) {
	visualizationID := c.Param("visualizationID")
	projectID := c.Param("projectID")
	if projectID == "" || visualizationID == "" {
		c.JSON(400, gin.H{"error": "you must specify a projectID, visualizationID, and version"})
		return
	}
	s3 := files.GetFileManager()
	vizKey := "projects/" + projectID + "/visualizations/" + visualizationID + "/"
	versionFiles, err := s3.ListFilesRecursive(vizKey)
	if err != nil {
		logrus.WithError(err).Error("error listing versions")
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	logrus.Infof("Found %d versions for %s", len(versionFiles), vizKey)
	versions := make([]map[string]interface{}, len(versionFiles))
	for idx, ver := range versionFiles {
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
		viz := map[string]interface{}{}
		viz["version"] = verNum
		err = s3.ReadJSON(files.GetVisualizationVersionKey(projectID, visualizationID, int(verNum)), &viz)
		if err != nil {
			logrus.WithError(err).Errorf("error reading visualization")
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		versions[idx] = viz
	}
	sort.Slice(versions, func(i, j int) bool {
		return versions[i]["version"].(int64) < versions[j]["version"].(int64)
	})
	c.JSON(200, versions)
}
