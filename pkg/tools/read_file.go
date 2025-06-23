package tools

import (
	"fmt"
	"os"

	"google.golang.org/genai"
)

const readFileDescription = `Read the contents of a file at a given relative ` +
	`file path named by the "path" argument. Use this to see the text inside ` +
	`the file. Do not use this with directory names.`

var readFileTool = &genai.Tool{
	FunctionDeclarations: []*genai.FunctionDeclaration{
		{
			Description: readFileDescription,
			Name:        "read_file",
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

func ReadFile(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read file '%s': %w", filePath, err)
	}
	return string(data), nil
}
