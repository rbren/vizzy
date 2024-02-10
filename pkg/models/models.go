package data

type DataUpload struct {
	UUID    string
	Key     string
	RawData []byte
}

type Metadata struct {
	Format      string
	Title       string
	Description string
	Fields      []string
}

type Visualization struct {
	UUID           string
	Details        string
	Visualizations map[string]VisualizationDetail
}

type VisualizationDetail struct {
	Prompt string
	Code   string
}
