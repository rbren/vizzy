package server

import (
	"os"
	"strings"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()

	r.Use(static.Serve("/", static.LocalFile("./app/dist", false)))
	r.NoRoute(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/api") {
			c.AbortWithStatus(404)
			return
		}
		c.File("./app/dist/index.html")
	})

	api := r.Group("/api")
	api.POST("/projects", uploadData)
	api.POST("/email", saveEmail)

	projectViewGroup := api.Group("/projects/:projectID")
	projectViewGroup.GET("/metadata", getMetadata)
	projectViewGroup.GET("/visualizations", listVisualizations)
	projectViewGroup.GET("/data", getData)
	projectViewGroup.GET("/fields-code", getFieldsCode)
	projectViewGroup.POST("/fork", forkProject)

	visualizationViewGroup := projectViewGroup.Group("/visualizations/:visualizationID")
	visualizationViewGroup.GET("/versions/:version", getVisualization)
	visualizationViewGroup.GET("/latest", getLatestVisualization)
	visualizationViewGroup.GET("/versions", listVersions)

	projectEditGroup := api.Group("/projects/:projectID")
	projectEditGroup.Use(checkKey)

	projectEditGroup.POST("/metadata", setMetadata)
	projectEditGroup.POST("/analyze", analyzeData)
	projectEditGroup.POST("/visualizations", createVisualization)
	projectEditGroup.POST("/fields-code", describeFields)
	projectEditGroup.POST("/fields-metadata", setFieldsMetadata)
	projectEditGroup.DELETE("", deleteProject)

	visualizationEditGroup := projectEditGroup.Group("/visualizations/:visualizationID")
	visualizationEditGroup.DELETE("", deleteVisualization)
	visualizationEditGroup.PATCH("/versions/:versionID", updateVisualization)
	visualizationEditGroup.DELETE("/versions/:versionID", deleteVisualizationVersion)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3031"
	}
	r.Run("0.0.0.0:" + port)
}
