package tools

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestToolsIntegration(t *testing.T) {
	t.Run("CreateFile", testCreateFile)
	t.Run("EditFile", testEditFile)
	t.Run("ListFiles", testListFiles)
	t.Run("ReadFile", testReadFile)
}

func testCreateFile(t *testing.T) {
	t.Parallel()
	tempDir, err := os.MkdirTemp("", "test-create-file")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	filePath := filepath.Join(tempDir, "newfile.txt")
	content := "Hello, world!"

	err = CreateFileWithDefaults(filePath, content, nil)
	if err != nil {
		t.Fatalf("CreateFile failed: %v", err)
	}

	actualContent, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	if string(actualContent) != content {
		t.Errorf("File content mismatch: got %q, want %q", actualContent, content)
	}
}

func testEditFile(t *testing.T) {
	t.Parallel()
	tempDir, err := os.MkdirTemp("", "test-edit-file")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	filePath := filepath.Join(tempDir, "existingfile.txt")
	initialContent := "This is the initial content."
	err = os.WriteFile(filePath, []byte(initialContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create initial file: %v", err)
	}

	oldString := "initial"
	newString := "updated"

	err = EditFile(filePath, oldString, newString)
	if err != nil {
		t.Fatalf("EditFile failed: %v", err)
	}

	actualContent, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	expectedContent := "This is the updated content."
	if string(actualContent) != expectedContent {
		t.Errorf("File content mismatch: got %q, want %q", actualContent, expectedContent)
	}
}

func testListFiles(t *testing.T) {
	t.Parallel()
	tempDir, err := os.MkdirTemp("", "test-list-files")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	err = os.WriteFile(filepath.Join(tempDir, "file1.txt"), []byte("content1"), 0644)
	if err != nil {
		t.Fatalf("Failed to create file1: %v", err)
	}
	err = os.MkdirAll(filepath.Join(tempDir, "subdir"), 0755)
	if err != nil {
		t.Fatalf("Failed to create subdir: %v", err)
	}
	err = os.WriteFile(filepath.Join(tempDir, "subdir", "file2.txt"), []byte("content2"), 0644)
	if err != nil {
		t.Fatalf("Failed to create file2: %v", err)
	}

	fileList, err := ListFiles(tempDir)
	if err != nil {
		t.Fatalf("ListFiles failed: %v", err)
	}

	expectedFiles := []string{"file1.txt", "subdir/", "subdir/file2.txt"}

	// Unmarshal the json string to a string array
	var actualFiles []string
	err = json.Unmarshal([]byte(fileList), &actualFiles)
	if err != nil {
		t.Fatalf("Failed to unmarshal file list: %v", err)
	}

	if !cmp.Equal(actualFiles, expectedFiles) {
		t.Errorf("File list mismatch: got %v, want %v", actualFiles, expectedFiles)
	}
}

func testReadFile(t *testing.T) {
	t.Parallel()
	tempDir, err := os.MkdirTemp("", "test-read-file")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	filePath := filepath.Join(tempDir, "readfile.txt")
	expectedContent := "This is the content to be read."
	err = os.WriteFile(filePath, []byte(expectedContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}

	content, err := ReadFile(filePath)
	if err != nil {
		t.Fatalf("ReadFile failed: %v", err)
	}

	if content != expectedContent {
		t.Errorf("File content mismatch: got %q, want %q", content, expectedContent)
	}
}
