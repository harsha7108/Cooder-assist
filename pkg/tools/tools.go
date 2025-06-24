// Package tools provides a set of tools that can be used by the AI.
package tools

import (
	"fmt"
	"iter"
	"maps"
	"slices"

	"google.golang.org/genai"
)

type Tools []*genai.Tool // Tools is a slice of genai.Tool.

// New creates a new set of tools.
func New() Tools {
	genAITools := make([]*genai.Tool, 0)
	genAITools = append(genAITools, readFileTool, listFilesTool, editFileTool, createFileTool, gitCommitTool, diffTool)
	return genAITools
}

var (
	// expectedArgs defines the expected arguments for each tool.
	expectedArgs = map[string]iter.Seq[string]{
		"read_file":   maps.Keys(readFileTool.FunctionDeclarations[0].Parameters.Properties),
		"list_files":  maps.Keys(listFilesTool.FunctionDeclarations[0].Parameters.Properties),
		"edit_file":   maps.Keys(editFileTool.FunctionDeclarations[0].Parameters.Properties),
		"create_file": maps.Keys(createFileTool.FunctionDeclarations[0].Parameters.Properties),
		"git_commit":  maps.Keys(gitCommitTool.FunctionDeclarations[0].Parameters.Properties),
		"diff":        maps.Keys(diffTool.FunctionDeclarations[0].Parameters.Properties),
	}
)

// ExecuteTool executes the tool specified in the given genai.FunctionCall.
// It validates the arguments and calls the appropriate tool function.
func (t Tools) ExecuteTool(call *genai.FunctionCall) *genai.FunctionResponse {
	response := make(map[string]any)

	name := call.Name
	// Get the arguments passed to the tool
	received := maps.Keys(call.Args)
	// Check if the tool name is valid
	if _, ok := expectedArgs[name]; !ok {
		response["error"] = fmt.Errorf("tool named '%s' is unknown", name)
	} else {
		// Check if all required parameters are present
		var missingRequired []string
		var requiredParams []string

		// Get required parameters based on tool name
		switch name {
		case "read_file", "list_files":
			requiredParams = []string{"path"}
		case "edit_file":
			requiredParams = []string{"path", "old_string", "new_string"}
		case "create_file":
			requiredParams = []string{"path", "content"}
		case "git_commit":
			requiredParams = []string{"message"}
		case "diff":
			requiredParams = []string{"old_string", "new_string"}
		}

		// Check for missing required parameters
		receivedSlice := slices.Collect(received)
		for _, required := range requiredParams {
			if !slices.Contains(receivedSlice, required) {
				missingRequired = append(missingRequired, required)
			}
		}

		if len(missingRequired) > 0 {
			response["error"] = fmt.Errorf("for tool named '%s' missing required arguments: %v", name, missingRequired)
		} else if call.Name == "read_file" {
			// Execute read file tool
			content, err := ReadFile(call.Args["path"].(string))
			if err == nil {
				response["output"] = content
			} else {
				response["error"] = err.Error()
			}
		} else if call.Name == "list_files" {
			// Execute list files tool
			content, err := ListFiles(call.Args["path"].(string))
			if err == nil {
				response["output"] = content
			} else {
				response["error"] = err.Error()
			}
		} else if call.Name == "edit_file" {
			// Execute edit file tool
			err := EditFile(call.Args["path"].(string), call.Args["old_string"].(string), call.Args["new_string"].(string))
			if err == nil {
				response["output"] = "OK"
			} else {
				response["error"] = err.Error()
			}
		} else if call.Name == "create_file" {
			// Execute create file tool
			path := call.Args["path"].(string)
			content := call.Args["content"].(string)

			// Handle optional overwrite parameter
			var overwrite *bool
			if overwriteVal, exists := call.Args["overwrite"]; exists {
				if boolVal, ok := overwriteVal.(bool); ok {
					overwrite = &boolVal
				}
			}

			err := CreateFileWithDefaults(path, content, overwrite)
			if err == nil {
				response["output"] = fmt.Sprintf("File '%s' created successfully", path)
			} else {
				response["error"] = err.Error()
			}
		} else if call.Name == "git_commit" {
			message := call.Args["message"].(string)
			content, err := GitCommit(message)
			if err == nil {
				response["output"] = content
			} else {
				response["error"] = err.Error()
			}
		} else if call.Name == "diff" {
			// Execute diff tool
			content, err := Diff(call.Args["old_string"].(string), call.Args["new_string"].(string))
			if err == nil {
				response["output"] = content
			} else {
				response["error"] = err.Error()
			}

		}
	}

	return &genai.FunctionResponse{
		ID:       call.ID,
		Name:     call.Name,
		Response: response,
	}
}
