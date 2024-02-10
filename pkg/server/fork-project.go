package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/rbren/vizzy/pkg/files"
)

func forkProject(c *gin.Context) {
	s3 := files.GetFileManager()
	sourceProjectID := c.Param("projectID")
	newProjectID := files.GenerateUUID()
	logrus.Infof("Forking project %s to %s", sourceProjectID, newProjectID)

	var metadata map[string]interface{}
	err := s3.ReadJSON(files.GetMetadataKey(sourceProjectID), &metadata)
	if err != nil {
		logrus.WithError(err).Errorf("error getting metatdata from s3 for project %s", sourceProjectID)
		c.JSON(http.StatusNotFound, gin.H{"error": "project metadata not found"})
		return
	}

	err = s3.CopyDirectory(files.GetProjectDirectoryKey(sourceProjectID), files.GetProjectDirectoryKey(newProjectID))
	if err != nil {
		logrus.WithError(err).Error("Failed to copy project")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to copy project"})
		return
	}
	authKey := files.GenerateUUID()
	err = s3.WriteFile(files.GetAccessKeyKey(newProjectID), []byte(authKey))
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate authentication key"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"uuid": newProjectID, "key": authKey})
}
