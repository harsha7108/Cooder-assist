// Package tools provides utility functions for interacting with the file system.
package tools

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/genai"
)

const createFileDescription = `Create a new file at the given 'path' with the specified 'content'. ` + // Description of the create file tool for the Gemini API
	`If the file already exists, it will be overwritten unless 'overwrite' is set to false. ` +
	`Directory structure will be created automatically if it doesn't exist.`

var createFileTool = &genai.Tool{ // Definition of the create file tool
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

func CreateFile(filePath, content string, overwrite bool) error { // CreateFile creates a new file with the given content and overwrite option.
	if filePath == "" {
		return errors.New("invalid argument: file path is empty") // Return error if file path is empty
	}

	// Clean the file path
	filePath = filepath.Clean(filePath)

	// Check if file already exists
	if _, err := os.Stat(filePath); err == nil {
		// File exists
		if !overwrite {
			return fmt.Errorf("file %s already exists and overwrite is disabled", filePath) // Return error if file exists and overwrite is disabled
		}
	} else if !os.IsNotExist(err) {
		// Some other error occurred while checking file
		return fmt.Errorf("failed to check if file exists %s: %w", filePath, err) // Return error if some other error occurred while checking file
	}

	// Create directory structure if it doesn't exist
	dir := filepath.Dir(filePath)
	if dir != "." && dir != "" {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err) // Return error if failed to create directory
		}
	}

	// Create/write the file
	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", filePath, err) // Return error if failed to create file
	}

	return nil // Return nil if file creation is successful
}

// Wrapper function that handles the default overwrite parameter
func CreateFileWithDefaults(filePath, content string, overwrite *bool) error { // CreateFileWithDefaults creates a new file with default overwrite parameter.
	// Default overwrite to true if not specified
	shouldOverwrite := true
	if overwrite != nil {
		shouldOverwrite = *overwrite
	}

	return CreateFile(filePath, content, shouldOverwrite) // Call CreateFile with the determined overwrite value
}
