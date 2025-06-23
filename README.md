# Cooder Assist

## Overview

Cooder Assist is a command-line tool designed to assist developers with coding-related tasks. It leverages the Gemini API to provide intelligent suggestions, file manipulation capabilities, and Git integration.

## Features

*   **Intelligent Code Suggestions:** Uses the Gemini API to provide context-aware code suggestions and answers to coding questions.
*   **File System Interaction:**
    *   Create new files with specified content.
    *   Edit existing files by replacing strings.
    *   List files and directories in a given path.
    *   Read the content of files.
    *   Determine the file type.
*   **Git Integration:** Stages all changes and creates a new Git commit with a provided message.
*   **Configuration:** Uses a YAML configuration file to manage client settings (e.g., Gemini model).
*   **Logging:** Provides logging functionality for debugging and monitoring.

## Usage

Before using this tool, you need to export the Google API key as an environment variable:

```bash
export GOOGLE_API_KEY="YOUR_API_KEY"
```



1.  **Configure the agent:** Modify the `agent-config-default.yml` file to set the desired Gemini model and other client settings.
2.  **Build the application:** Run `make` to build the binaries.
3.  **Run the application:** Execute the `cooder-assist-local` binary from the `bin/` directory.


## File Descriptions

### Configuration Files

*   **`agent-config-default.yml`**: Default configuration file for the agent, specifying client settings like the Gemini model.
    ```yaml
    Client:
      Name: "Gemini"
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

The `Makefile` provides targets for building the application.

*   **`make build`**: Builds the `cooder-assist-local` binary for your local operating system and architecture.
*   **`make buildlinux`**: Builds the `cooder-assist` binary specifically for Linux (amd64 architecture).
*   **`make default`**: Executes `fmt`, `lint`, `buildlinux`, and `build` in sequence.

**Building the code:**

To build the application, run:

```bash
make
```

This will generate the `cooder-assist-local` and `cooder-assist` binaries in the `bin/` directory.

## Dependencies

The application uses the following Go modules:

```
cloud.google.com/go v0.116.0
cloud.google.com/go/auth v0.13.0
cloud.google.com/go/compute/metadata v0.6.0
github.com/felixge/httpsnoop v1.0.4
github.com/fsnotify/fsnotify v1.8.0
github.com/go-logr/logr v1.4.2
github.com/go-logr/stdr v1.2.2
github.com/go-viper/mapstructure/v2 v2.2.1
github.com/google/go-cmp v0.6.0
github.com/google/s2a-go v0.1.8
github.com/googleapis/enterprise-certificate-proxy v0.3.4
github.com/googleapis/gax-go/v2 v2.14.1
github.com/gorilla/websocket v1.5.3
github.com/inconshreveable/mousetrap v1.1.0
github.com/pelletier/go-toml/v2 v2.2.3
github.com/sagikazarmark/locafero v0.7.0
github.com/sourcegraph/conc v0.3.0
github.com/spf13/afero v1.12.0
github.com/spf13/cast v1.7.1
github.com/spf13/cobra v1.9.1
github.com/spf13/pflag v1.0.6
github.com/spf13/viper v1.20.1
github.com/subosito/gotenv v1.6.0
go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.54.0
go.opentelemetry.io/otel v1.29.0
go.opentelemetry.io/otel/metric v1.29.0
go.opentelemetry.io/otel/trace v1.29.0
go.uber.org/atomic v1.9.0
go.uber.org/multierr v1.9.0
golang.org/x/crypto v0.32.0
golang.org/x/net v0.33.0
golang.org/x/sys v0.29.0
golang.org/x/text v0.21.0
google.golang.org/genai v1.12.0
google.golang.org/genproto/googleapis/rpc v0.0.0-20241223144023-3abc09e42ca8
google.golang.org/grpc v1.67.3
google.golang.org/protobuf v1.36.1
gopkg.in/yaml.v3 v3.0.1
```



## Contributing

Contributions are welcome! Please submit a pull request with your changes.
