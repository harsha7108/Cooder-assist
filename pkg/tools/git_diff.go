// This file implements the git diff functionality.
// It allows displaying the git diff of a specified file.
package tools

import (
	"fmt"
	"os/exec"

	"google.golang.org/genai"
)

const gitDiffDescription = `Display the gitdiff of the file ` +
	`execute this tool before executing edit_file and` +
	`ask for approval for edit file after showing the diff`

var gitDiffTool = &genai.Tool{
	FunctionDeclarations: []*genai.FunctionDeclaration{
		{
			Description: gitDiffDescription,
			Name:        "git_diff",
			Parameters: &genai.Schema{
				Type: genai.TypeObject,
				Properties: map[string]*genai.Schema{
					"path": {
						Type: genai.TypeString,
					}},
			},
		},
	},
}

func GitDiff(path string) (string, error) {
	cmd := exec.Command("git", "diff", path)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to stage changes: %w, output: %s", err, string(output))
	}
	return string(output), nil
}
