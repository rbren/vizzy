package server

import (
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/rbren/go-prompter/pkg/files"

	"github.com/rbren/vizzy/pkg/keys"
)

func saveEmail(c *gin.Context) {
	s3 := files.GetFileManager()

	email := c.Query("email")
	_, err := mail.ParseAddress(email)
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email"})
		return
	}

	err = s3.WriteFile(keys.GetEmailKey(email), []byte(email))
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}
