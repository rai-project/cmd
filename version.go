package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

type VersionInfo struct {
	// Version is populated at compile time by govvv from ./VERSION
	Version string
	// GitCommit is populated at compile time by govvv.
	BuildDate string
	// GitState is populated at compile time by govvv.
	GitCommit  string
	GitBranch  string
	GitState   string
	GitSummary string
}

var Version = VersionInfo{
	Version:    "0.2.0",
	BuildDate:  time.Now().String(),
	GitCommit:  "-dirty-",
	GitBranch:  "-dirty-",
	GitState:   "-dirty-",
	GitSummary: "-dirty-",
}

// versionCmd represents the version command
var VersionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"buildinfo"},
	Short:   "Prints a the version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("BuildDate: ", Version.BuildDate)
		fmt.Println("GitCommit: ", Version.GitCommit)
		fmt.Println("GitBranch: ", Version.GitBranch)
		fmt.Println("GitState: ", Version.GitState)
		fmt.Println("GitSummary: ", Version.GitSummary)
	},
}
