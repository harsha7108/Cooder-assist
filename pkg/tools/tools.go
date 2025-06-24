package tools

import (
	"fmt"
	"iter"
	"maps"
	"slices"

	"google.golang.org/genai"
)

type Tools []*genai.Tool

func New() Tools {
	genAITools := make([]*genai.Tool, 0)
	genAITools = append(genAITools, readFileTool, listFilesTool, editFileTool, createFileTool, gitCommitTool, gitDiffTool)
	return genAITools
}

var (
	expectedArgs = map[string]iter.Seq[string]{
		"read_file":   maps.Keys(readFileTool.FunctionDeclarations[0].Parameters.Properties),
		"list_files":  maps.Keys(listFilesTool.FunctionDeclarations[0].Parameters.Properties),
		"edit_file":   maps.Keys(editFileTool.FunctionDeclarations[0].Parameters.Properties),
		"create_file": maps.Keys(createFileTool.FunctionDeclarations[0].Parameters.Properties),
		"git_commit":  maps.Keys(gitCommitTool.FunctionDeclarations[0].Parameters.Properties),
		"git_diff":    maps.Keys(gitDiffTool.FunctionDeclarations[0].Parameters.Properties),
	}
)

func (t Tools) ExecuteTool(call *genai.FunctionCall) *genai.FunctionResponse {
	response := make(map[string]any)

	name := call.Name
	received := maps.Keys(call.Args)
	if _, ok := expectedArgs[name]; !ok {
		response["error"] = fmt.Errorf("tool named '%s' is unknown", name)
	} else {
		// Check if all required parameters are present
		var missingRequired []string
		var requiredParams []string

		// Get required parameters based on tool name
		switch name {
		case "read_file", "list_files", "git_diff":
			requiredParams = []string{"path"}
		case "edit_file":
			requiredParams = []string{"path", "old_string", "new_string"}
		case "create_file":
			requiredParams = []string{"path", "content"}
		case "git_commit":
			requiredParams = []string{"message"}
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
			content, err := ReadFile(call.Args["path"].(string))
			if err == nil {
				response["output"] = content
			} else {
				response["error"] = err.Error()
			}
		} else if call.Name == "list_files" {
			content, err := ListFiles(call.Args["path"].(string))
			if err == nil {
				response["output"] = content
			} else {
				response["error"] = err.Error()
			}
		} else if call.Name == "edit_file" {
			err := EditFile(call.Args["path"].(string), call.Args["old_string"].(string), call.Args["new_string"].(string))
			if err == nil {
				response["output"] = "OK"
			} else {
				response["error"] = err.Error()
			}
		} else if call.Name == "create_file" {
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
		} else if call.Name == "git_diff" {
			content, err := GitDiff(call.Args["path"].(string))
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
