package cmd

import (
	"fmt"

	"github.com/rai-project/config"
	"github.com/spf13/cobra"
)

var Version = &config.App.Version

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
