# Cooder Assist

## Overview

Cooder Assist is a command-line tool designed to assist developers with coding-related tasks. It leverages the Gemini API to provide intelligent suggestions, file manipulation capabilities, and Git integration.

## Features

*   **Intelligent Code Suggestions:** Uses the Gemini API to provide context-aware code suggestions and answers to coding questions.
*   **Git Integration:** Stages all changes and creates a new Git commit with a provided message.
*   **Configuration:** Uses a YAML configuration file to manage client settings (e.g., Gemini model).
*   **Logging:** Provides logging functionality for debugging and monitoring.

## Usage

Before using this tool, you need to export the Google API key as an environment variable:

```bash
export GOOGLE_API_KEY="YOUR_API_KEY"
```



**Configure the agent:** Modify the `agent-config-default.yml` file or write your own config to set the desired Gemini model. 
**Build the application:** 
The `Makefile` provides targets for building the application.To build the application follow the instructions below

*   **`make build`**: Builds the `cooder-assist-local` binary for your local operating system and architecture.
*   **`make buildlinux`**: Builds the `cooder-assist` binary specifically for Linux (amd64 architecture).

This will generate the `cooder-assist-local` and `cooder-assist` binaries in the `bin/` directory.
**Running the Application:**

To run the application, move the binary to the root of your codebase and run :

```bash
./cooder-assist-local --cfgPath=. --cfgFile=agent-config-default.yml
```
Where cfgPath is the path to config and cfgFile is the file name of the config


## File Descriptions

### Configuration Files

*   **`agent-config-default.yml`**: Default configuration file for the agent, specifying client settings like the Gemini model.
    ```yaml
    ModelConfig:
      Model: "gemini-2.0-flash"
    ```

### Go Files

*   **`cmd/main.go`**: The main entry point of the `cooder-assist-local` application. It initializes and executes the root command.
*   **`cmd/provisioner.go`**: This file contains the `newProvisioner` function, which sets up the Gemini client, configures tools, and starts the agent.
*   **`pkg/agent/gemini.go`**: Defines the `Agent` struct and its methods for interacting with the Gemini API. It handles user input, sends messages to Gemini, and processes the responses.
*   **`pkg/config/config.go`**: Handles the configuration loading and management for the application. It reads the configuration from a YAML file and provides access to the configuration values.
*   **`pkg/log/logger.go`**: Implements the logging functionality for the application using the `slog` package.
*   **`pkg/scanner/scanner.go`**: Provides a scanner for reading user input from the command line.
*   **`pkg/tools/create_file.go`**: Implements the `create_file` tool, which creates a new file with specified content.
*   **`pkg/tools/edit_file.go`**: Implements the `edit_file` tool, which edits an existing file by replacing instances of a string with another string.
*   **`pkg/tools/get_file_type.go`**: Implements the `get_file_type` tool, which determines the file type of a given file.
*   **`pkg/tools/git_commit.go`**: Implements the `git_commit` tool, which stages all changes and creates a new Git commit with the given message.
*   **`pkg/tools/list_files.go`**: Implements the `list_files` tool, which lists files and directories in a given path.
*   **`pkg/tools/read_file.go`**: Implements the `read_file` tool, which reads the content of a file.
*   **`pkg/tools/tools.go`**: Defines the `Tools` type and provides methods for managing and executing the available tools.
*   **`pkg/tools/tools_integration_test.go`**: Contains integration tests for the tools package.

### Makefile








## Contributing

Contributions are welcome! Please submit a pull request with your changes.
