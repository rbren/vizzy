package tcase

import (
	"strings"
)

type promptSet struct {
	base      string
	followups []string
}

func parsePrompts(s string) []promptSet {
	prompts := []promptSet{}
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		if strings.HasPrefix(line, "  ") {
			line = strings.TrimSpace(line)
			curPrompt := &prompts[len(prompts)-1]
			curPrompt.followups = append(curPrompt.followups, line)
		} else {
			prompts = append(prompts, promptSet{
				base:      line,
				followups: []string{},
			})
		}
	}
	return prompts
}
