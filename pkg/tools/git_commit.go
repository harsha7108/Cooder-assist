package tools

import (
	"fmt"
	"os/exec"

	"google.golang.org/genai"
)

const gitCommitDescription = `Stages all changes and creates a new Git commit with the given message. DONT COMMIT WITHOUT APPROVAL. ALWAYS ASK THE USER FOR APPROVAL`

var gitCommitTool = &genai.Tool{
	FunctionDeclarations: []*genai.FunctionDeclaration{
		{
			Description: gitCommitDescription,
			Name:        "git_commit",
			Parameters: &genai.Schema{
				Type: genai.TypeObject,
				Properties: map[string]*genai.Schema{
					"message": {
						Type:        genai.TypeString,
						Description: "The commit message.",
					},
				},
				Required: []string{"message"},
			},
		},
	},
}

func GitCommit(message string) (string, error) {
	// Stage all changes
	cmd := exec.Command("git", "add", ".")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to stage changes: %w, output: %s", err, string(output))
	}

	// Commit changes
	cmd = exec.Command("git", "commit", "-m", message)
	output, err = cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to commit changes: %w, output: %s", err, string(output))
	}

	return "Commit successful", nil
}
