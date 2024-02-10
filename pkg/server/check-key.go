package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/rbren/vizzy/pkg/files"
)

func checkKey(c *gin.Context) {
	projectID := c.Param("projectID")
	keyProvided := c.GetHeader("X-PROJECT-KEY")
	s3 := files.GetFileManager()

	actualKey, err := s3.ReadFile(files.GetAccessKeyKey(projectID))
	if err != nil {
		logrus.WithError(err).Error("error reading key file")
		c.JSON(http.StatusNotFound, gin.H{"error": "project not found"})
		c.Abort()
		return
	}

	if keyProvided != string(actualKey) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		c.Abort()
		return
	}
	c.Next()
}
