package tcase

import (
	"embed"
	"html/template"
	"os"
)

//go:embed template.html
var templateFS embed.FS

// VisualizationTemplateData contains the data needed to fill the template.
type VisualizationTemplateData struct {
	Code     string
	DataURL  string
	Title    string
	Prompt   string
	TestCase string
	Group    int
	Update   int
}

// writeHTML fills the given template with data and writes the output to a file.
func writeHTML(outputPath string, data VisualizationTemplateData) error {
	tmpl, err := template.ParseFS(templateFS, "template.html")
	if err != nil {
		return err
	}

	input := map[string]interface{}{
		"Code":     template.HTML("<script>" + data.Code + "</script>"),
		"DataURL":  template.HTML(data.DataURL),
		"TestCase": data.TestCase,
		"Title":    data.Title,
		"Prompt":   data.Prompt,
		"Group":    data.Group,
		"Update":   data.Update,
	}

	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	err = tmpl.Execute(outputFile, input)
	if err != nil {
		return err
	}

	return nil
}
