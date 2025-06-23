package tools

import (
	"fmt"
	"net/http"
	"os"

	"google.golang.org/genai"
)

const getFileTypeDescription = `Determine the file type of a file.  It uses the file extension and reads the first 512 bytes to determine the file type.`

var getFileTypeTool = &genai.Tool{
	FunctionDeclarations: []*genai.FunctionDeclaration{
		{
			Description: getFileTypeDescription,
			Name:        "get_file_type",
			Parameters: &genai.Schema{
				Type: genai.TypeObject,
				Properties: map[string]*genai.Schema{
					"path": {
						Type:        genai.TypeString,
						Description: "The path to the file.",
					},
				},
				Required: []string{"path"},
			},
		},
	},
}

func GetFileType(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Read the first 512 bytes of the file
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil && err.Error() != "EOF" {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	// Get the file type based on the content
	fileType := http.DetectContentType(buffer)

	return fileType, nil
}
