// This file implements the git diff functionality.
// It allows displaying the git diff of a specified file.
package tools

import (
	"fmt"
	"os/exec"

	"google.golang.org/genai"
)

const gitDiffDescription = `Display the git diff of the file. This tool should be executed to review the changes before using the edit_file tool.  The user is responsible for calling this tool and reviewing the output before calling edit_file.`

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
		exitError, ok := err.(*exec.ExitError)
		if ok {
			return "", fmt.Errorf("git diff command failed with exit code %d, output: %s", exitError.ExitCode(), string(output))

		}
		return "", fmt.Errorf("git diff command failed: %w, output: %s", err, string(output))
	}
	return string(output), nil
}
