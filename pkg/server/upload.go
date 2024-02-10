package server

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/rbren/vizzy/pkg/files"
)

const maxFileSize = 100 << 20 // 100MB
const maxFileSizeString = "100MB"

func uploadData(c *gin.Context) {
	s3 := files.GetFileManager()

	var dataBody io.ReadCloser
	var dataLength int64

	url := c.Query("url")
	if url == "" {
		dataBody = http.MaxBytesReader(c.Writer, c.Request.Body, maxFileSize)
		dataLength = c.Request.ContentLength
	} else {
		resp, err := http.Get(url)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
			return
		}
		dataBody = http.MaxBytesReader(c.Writer, resp.Body, maxFileSize)
		dataLength = resp.ContentLength
	}
	defer dataBody.Close()

	if dataLength > maxFileSize {
		c.JSON(http.StatusRequestEntityTooLarge, gin.H{"error": "File size exceeds " + maxFileSizeString})
		return
	}

	rawData, err := ioutil.ReadAll(dataBody)
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusRequestEntityTooLarge, gin.H{"error": "Invalid data or data size too large"})
		return
	}

	projectID := files.GenerateUUID()

	err = s3.WriteFile(files.GetDataKey(projectID), rawData)
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload data"})
		return
	}

	authKey := files.GenerateUUID()
	err = s3.WriteFile(files.GetAccessKeyKey(projectID), []byte(authKey))
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload authentication key"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"uuid": projectID, "key": authKey})
}
