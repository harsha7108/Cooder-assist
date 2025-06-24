// This file implements the git diff functionality.
// It allows displaying the git diff of a specified file.
package tools

import (
	"fmt"
	"os"
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

	oldFile, err := os.CreateTemp("", "old-*.txt")
	if err != nil {
		return "", err
	}
	defer os.Remove(oldFile.Name())

	newFile, err := os.CreateTemp("", "new-*.txt")
	if err != nil {
		return "", err
	}
	defer os.Remove(newFile.Name())

	oldFile.WriteString(oldStr)
	newFile.WriteString(newStr)
	oldFile.Close()
	newFile.Close()

	cmd := exec.Command("diff", "-u", oldFile.Name(), newFile.Name())
	output, err := cmd.CombinedOutput()

	if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() == 1 {
		return string(output), nil
	}
	if err != nil {
		return "", fmt.Errorf("diff error: %w\noutput: %s", err, string(output))
	}
	return string(output), nil
}
