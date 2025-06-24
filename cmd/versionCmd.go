package main

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

// values set during build
var (
	version      = "" // git hash or version num passed in during build
	gitCommit    = ""
	gitTreeState = ""
	buildDate    = ""
	buildTime    = ""
)

var versionTemplate = `  Version:	%v
  GitCommit:	%v
  GitTreeState:	%v
  BuildDate:	%v
  BuildTime:	%v
  GoVersion:	%v
`

// BuildInfo describes the compile time information.
type BuildInfo struct {
	// Version is the current get hash or version passed in during build.
	Version string `json:"version"`
	// GitCommit is the git commit id.
	GitCommit string `json:"git_commit"`
	// GitTreeState is the state of the git tree.
	GitTreeState string `json:"git_tree_state"`

	GoVersion string `json:"go_version"`

	BuildDate string `json:"build_date"`

	BuildTime string `json:"build_time"`
}

func getVersion() string {
	v := BuildInfo{
		Version:      version,
		GitCommit:    gitCommit,
		GitTreeState: gitTreeState,
		BuildDate:    buildDate,
		BuildTime:    buildTime,
		GoVersion:    runtime.Version(),
	}

	return fmt.Sprintf(versionTemplate, v.Version, v.GitCommit, v.GitTreeState, v.BuildDate, v.BuildTime, v.GoVersion)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Provides version and build info",
	Long: `version provides build information like when the binary was built, 
what the state of the git tree was, what GO version was used to compile etc.`,
	Example: "provisioner version",
	Run: func(cmd *cobra.Command, args []string) {
		v := getVersion()
		fmt.Print(v)
	},
}
