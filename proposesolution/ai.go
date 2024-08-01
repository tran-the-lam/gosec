package proposesolution

import (
	"context"
	"fmt"
	"time"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"

	"github.com/securego/gosec/v2/issue"
)

const (
	GeminiModel    = "gemini-1.5-flash"
	AIPrompt       = "What is the solution to fix the error \"%s\". Answer limited to 200 words"
	GeminiProvider = "gemini"
)

func generateSolutionByGemini(aiApiKey string, issues []*issue.Issue) error {
	timeout := 30 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	client, err := genai.NewClient(ctx, option.WithAPIKey(aiApiKey))
	if err != nil {
		return err
	}
	defer client.Close()

	model := client.GenerativeModel(GeminiModel)
	for _, issue := range issues {
		prompt := fmt.Sprintf(AIPrompt, issue.What)
		resp, err := model.GenerateContent(ctx, genai.Text(prompt))
		if err != nil {
			return err
		}

		issue.ProposedSolution = fmt.Sprintf("%+v", resp.Candidates[0].Content.Parts[0])
	}
	return nil
}

// GenerateSolution generates a solution for the given issues using the specified AI provider
func GenerateSolution(aiApiProvider, aiApiKey string, issues []*issue.Issue) error {
	switch aiApiProvider {
	case GeminiProvider:
		return generateSolutionByGemini(aiApiKey, issues)
	default:
		return fmt.Errorf("ai provider not supported")
	}
}
