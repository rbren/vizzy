package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/rbren/vizzy/pkg/files"
)

func setFieldsMetadata(c *gin.Context) {
	projectID := c.Param("projectID")
	if projectID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Project ID is required"})
		return
	}
	s3 := files.GetFileManager()

	fieldsData := map[string]interface{}{}
	err := s3.ReadJSON(files.GetFieldsCodeKey(projectID), &fieldsData)
	if err != nil {
		logrus.WithError(err).Errorf("error reading fields code from s3 for project %s", projectID)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error retrieving fields code"})
		return
	}

	bodyData := map[string]interface{}{}
	if err := c.BindJSON(&bodyData); err != nil {
		logrus.WithError(err).Errorf("error binding json for project %s", projectID)
		c.JSON(http.StatusBadRequest, gin.H{"error": "error parsing request body"})
		return
	}

	fieldsData["metadata"] = bodyData
	err = s3.WriteJSON(files.GetFieldsCodeKey(projectID), fieldsData)
	if err != nil {
		logrus.WithError(err).Errorf("error writing fields code to s3 for project %s", projectID)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error writing fields metadata"})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"success": true})
}
