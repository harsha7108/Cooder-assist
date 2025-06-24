// This file implements the git diff functionality.
// It allows displaying the git diff of a specified file.
package tools

import (
	"fmt"
	"os/exec"

	"google.golang.org/genai"
)

const diffDescription = `Display the diff between two strings . This tool should be executed to review the changes before using the edit_file tool.`

var diffTool = &genai.Tool{
	FunctionDeclarations: []*genai.FunctionDeclaration{
		{
			Description: diffDescription,
			Name:        "diff",
			Parameters: &genai.Schema{
				Type: genai.TypeObject,
				Properties: map[string]*genai.Schema{
					"old_string": {
						Type: genai.TypeString,
					},
					"new_string": {
						Type: genai.TypeString,
					}},
			},
		},
	},
}

func Diff(oldStr, newStr string) (string, error) {
	cmd := exec.Command("diff", "<(echo %s)", "<(echo %s)", oldStr, newStr)
	output, err := cmd.CombinedOutput()
	if err != nil {
		exitError, ok := err.(*exec.ExitError)
		if ok {
			return "", fmt.Errorf("diff command failed with exit code %d, output: %s", exitError.ExitCode(), string(output))

		}
		return "", fmt.Errorf("diff command failed: %w, output: %s", err, string(output))
	}
	return string(output), nil
}
