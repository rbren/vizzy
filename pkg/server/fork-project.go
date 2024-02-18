package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rbren/go-prompter/pkg/files"
	"github.com/sirupsen/logrus"

	"github.com/rbren/vizzy/pkg/keys"
)

func forkProject(c *gin.Context) {
	s3 := files.GetFileManager()
	sourceProjectID := c.Param("projectID")
	newProjectID := keys.GenerateUUID()
	logrus.Infof("Forking project %s to %s", sourceProjectID, newProjectID)

	var metadata map[string]interface{}
	err := s3.ReadJSON(keys.GetMetadataKey(sourceProjectID), &metadata)
	if err != nil {
		logrus.WithError(err).Errorf("error getting metatdata from s3 for project %s", sourceProjectID)
		c.JSON(http.StatusNotFound, gin.H{"error": "project metadata not found"})
		return
	}

	err = s3.CopyDirectory(keys.GetProjectDirectoryKey(sourceProjectID), keys.GetProjectDirectoryKey(newProjectID))
	if err != nil {
		logrus.WithError(err).Error("Failed to copy project")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to copy project"})
		return
	}
	authKey := keys.GenerateUUID()
	err = s3.WriteFile(keys.GetAccessKeyKey(newProjectID), []byte(authKey))
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate authentication key"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"uuid": newProjectID, "key": authKey})
}
