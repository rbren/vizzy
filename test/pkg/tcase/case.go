package tcase

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"

	"github.com/rbren/vizzy/pkg/files"
	"github.com/rbren/vizzy/pkg/llm"
	"github.com/rbren/vizzy/pkg/query"
)

const host = "http://localhost:3031"

const baseDir = "test/e2e/cases/"

func RunTestCase(name string) error {
	if os.Getenv("TEST_CASE") != "" && !strings.Contains(os.Getenv("TEST_CASE"), name) {
		logrus.Infof("Skip     %s", name)
		return nil
	}
	logrus.Infof("Starting %s", name)
	var llmClient llm.Client
	if os.Getenv("LLM_BACKEND") == "HUGGING_FACE" {
		logrus.Infof("Using Hugging Face API")
		llmClient = llm.NewHuggingFaceClient(os.Getenv("HUGGING_FACE_API_KEY"), os.Getenv("HUGGING_FACE_URL"))
	} else if os.Getenv("LLM_BACKEND") == "GEMINI" {
		gClient := llm.NewGeminiClient(os.Getenv("GEMINI_API_KEY"))
		llmClient = gClient
	} else {
		oaiClient := llm.NewOpenAIClient(os.Getenv("OPENAI_API_KEY"), os.Getenv("OPENAI_MODEL"))
		oaiClient.Seed = 42
		llmClient = oaiClient
	}
	llmClient.SetDebugFileManager(&files.LocalFileManager{
		BasePath: baseDir + "/" + name + "/debug/",
	})
	queryEngine := &query.Engine{llmClient}
	desc, err := analyzeData(name, queryEngine)
	if err != nil {
		logrus.Errorf("failed to analyze data: %v", err)
		return err
	}
	err = generateVisualizations(name, queryEngine, desc)
	if err != nil {
		logrus.Errorf("failed to generate visualizations: %v", err)
		return err
	}
	return nil
}

func generateVisualization(idx1, idx2 int, prompt, name string, queryEngine *query.Engine, desc *query.DataDescription, data string, prev *query.Visualization) (*query.Visualization, error) {
	debugDir := baseDir + "/" + name + "/debug/visualizations/" + fmt.Sprintf("%d/%d/", idx1, idx2)
	tmpOpenaiClient := query.Engine{queryEngine.LLM.Copy()}
	tmpOpenaiClient.LLM.SetDebugFileManager(&files.LocalFileManager{
		BasePath: debugDir,
	})
	logrus.Infof("Generating visualization with prompt: %s", prompt)
	vis, err := tmpOpenaiClient.CreateVisualization(prompt, *desc, nil, data, prev)
	if err != nil {
		return nil, err
	}
	err = writeHTML(debugDir+"index.html", VisualizationTemplateData{
		Code:     vis.Code,
		DataURL:  "/" + name + "/data",
		TestCase: name,
		Group:    idx1,
		Update:   idx2,
		Title:    vis.Title,
		Prompt:   prompt,
	})
	if err != nil {
		return nil, err
	}
	return vis, nil
}

func generateVisualizations(name string, queryEngine *query.Engine, desc *query.DataDescription) error {
	dir := baseDir + "/" + name
	b, err := os.ReadFile(dir + "/prompts.txt")
	if err != nil {
		return err
	}
	prompts := parsePrompts(string(b))

	b, err = os.ReadFile(dir + "/data")
	if err != nil {
		return err
	}
	data := string(b)

	var wg sync.WaitGroup
	for idx, prompt := range prompts {
		if os.Getenv("TEST_VISUALIZATION") != "" && strconv.Itoa(idx) != os.Getenv("TEST_VISUALIZATION") {
			logrus.Infof("Skip     %s", prompt.base)
			continue
		}
		wg.Add(1)
		go func(prompt promptSet, idx int) {
			defer wg.Done()
			vis, err := generateVisualization(idx, 0, prompt.base, name, queryEngine, desc, data, nil)
			if err != nil {
				logrus.Fatalf("failed to generate visualization %s: %v", prompt.base, err) // TODO: be more graceful
			}
			for followIdx, followupPrompt := range prompt.followups {
				if os.Getenv("TEST_SUBVISUALIZATION") != "" && strconv.Itoa(followIdx) != os.Getenv("TEST_SUBVISUALIZATION") {
					logrus.Infof("Skip     %s", followupPrompt)
					continue
				}
				vis, err = generateVisualization(idx, followIdx+1, followupPrompt, name, queryEngine, desc, data, vis)
				if err != nil {
					logrus.Fatalf("failed to generate visualization %s: %v", followupPrompt, err) // TODO: be more graceful
				}
			}
		}(prompt, idx)
	}
	wg.Wait()
	return nil
}

func analyzeData(name string, queryEngine *query.Engine) (*query.DataDescription, error) {
	b, err := os.ReadFile(baseDir + "/" + name + "/data")
	if err != nil {
		return nil, err
	}
	data := string(b)

	desc, err := queryEngine.DescribeData(data)
	if err != nil {
		return nil, err
	}

	return &desc, nil
}
