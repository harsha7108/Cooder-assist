package tools

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"google.golang.org/genai"
)

const listfilesDescription = "List files and directories at a given path. If no path is provided, lists files in the current directory."

var listFilesTool = &genai.Tool{
	FunctionDeclarations: []*genai.FunctionDeclaration{
		{
			Description: listfilesDescription,
			Name:        "list_files",
			Parameters: &genai.Schema{
				Type: genai.TypeObject,
				Properties: map[string]*genai.Schema{
					"path": {
						Type: genai.TypeString,
					},
				},
			},
		},
	},
}

func ListFiles(path string) (string, error) {
	if path == "" {
		path = "."
	}

	var files []string
	err := filepath.WalkDir(path, func(filePath string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("error walking directory: %w", err)
		}

		baseName, err := filepath.Rel(path, filePath)
		if err != nil {
			return fmt.Errorf("error getting relative path: %w", err)
		}

		if baseName != "." {
			if d.IsDir() {
				files = append(files, baseName+string(os.PathSeparator))
			} else {
				files = append(files, baseName)
			}
		}
		return nil
	})

	if err != nil {
		return "", err
	}

	result, err := json.Marshal(files)
	if err != nil {
		return "", fmt.Errorf("error marshaling JSON: %w", err)
	}
	return string(result), nil
}
