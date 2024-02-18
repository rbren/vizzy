package keys

import (
	"strconv"
)

func GetEmailKey(email string) string {
	return "emails/" + GenerateUUID() + ".txt"
}

func GetProjectDirectoryKey(projectID string) string {
	return "projects/" + projectID + "/"
}

func GetDataKey(projectID string) string {
	return GetProjectDirectoryKey(projectID) + "data"
}

func GetMetadataKey(projectID string) string {
	return GetProjectDirectoryKey(projectID) + "metadata.json"
}

func GetAccessKeyKey(projectID string) string {
	return GetProjectDirectoryKey(projectID) + "key.txt"
}

func GetFieldsCodeKey(projectID string) string {
	return GetProjectDirectoryKey(projectID) + "fields.json"
}

func GetVisualizationDirectoryKey(projectID string, visualizationID string) string {
	return GetProjectDirectoryKey(projectID) + "visualizations/" + visualizationID + "/"
}

func GetVisualizationVersionKey(projectID string, visualizationID string, version int) string {
	return GetVisualizationDirectoryKey(projectID, visualizationID) + strconv.Itoa(version) + ".json"
}
