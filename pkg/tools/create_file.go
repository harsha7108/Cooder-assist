package tools

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/genai"
)

const createFileDescription = `Create a new file at the given 'path' with the specified 'content'. ` +
	`If the file already exists, it will be overwritten unless 'overwrite' is set to false. ` +
	`Directory structure will be created automatically if it doesn't exist.`

var createFileTool = &genai.Tool{
	FunctionDeclarations: []*genai.FunctionDeclaration{
		{
			Description: createFileDescription,
			Name:        "create_file",
			Parameters: &genai.Schema{
				Type: genai.TypeObject,
				Properties: map[string]*genai.Schema{
					"path": {
						Type:        genai.TypeString,
						Description: "The file path where the new file should be created",
					},
					"content": {
						Type:        genai.TypeString,
						Description: "The content to write to the new file",
					},
					"overwrite": {
						Type:        genai.TypeBoolean,
						Description: "Whether to overwrite the file if it already exists (default: true)",
					},
				},
				Required: []string{"path", "content"},
			},
		},
	},
}

func CreateFile(filePath, content string, overwrite bool) error {
	if filePath == "" {
		return errors.New("invalid argument: file path is empty")
	}

	// Clean the file path
	filePath = filepath.Clean(filePath)

	// Check if file already exists
	if _, err := os.Stat(filePath); err == nil {
		// File exists
		if !overwrite {
			return fmt.Errorf("file %s already exists and overwrite is disabled", filePath)
		}
	} else if !os.IsNotExist(err) {
		// Some other error occurred while checking file
		return fmt.Errorf("failed to check if file exists %s: %w", filePath, err)
	}

	// Create directory structure if it doesn't exist
	dir := filepath.Dir(filePath)
	if dir != "." && dir != "" {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}

	// Create/write the file
	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", filePath, err)
	}

	return nil
}

// Wrapper function that handles the default overwrite parameter
func CreateFileWithDefaults(filePath, content string, overwrite *bool) error {
	// Default overwrite to true if not specified
	shouldOverwrite := true
	if overwrite != nil {
		shouldOverwrite = *overwrite
	}

	return CreateFile(filePath, content, shouldOverwrite)
}
