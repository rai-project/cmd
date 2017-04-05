package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// BuildTimeCmd represents the version command
var BuildTimeCmd = &cobra.Command{
	Use:   "buildtime",
	Short: "Prints a the time the binary was compiled",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("BuildDate: ", Version.BuildDate)
	},
}
