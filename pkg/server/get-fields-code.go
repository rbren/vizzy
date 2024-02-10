package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/rbren/vizzy/pkg/files"
)

func getFieldsCode(c *gin.Context) {
	s3 := files.GetFileManager()

	projectID := c.Param("projectID")
	if projectID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Project ID is required"})
		return
	}

	code := map[string]interface{}{}
	err := s3.ReadJSON(files.GetFieldsCodeKey(projectID), &code)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "fields code not found"})
		return
	}
	c.JSON(http.StatusOK, code)
}
