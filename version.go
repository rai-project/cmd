package cmd

import (
	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
)

type VersionInfo struct {
	BuildDate  string
	GitCommit  string
	GitBranch  string
	GitState   string
	GitSummary string
}

var Version VersionInfo

// versionCmd represents the version command
var VersionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"info", "v"},
	Short:   "Prints a the version information",
	Run: func(cmd *cobra.Command, args []string) {
		pp.Println(Version)
	},
}
