package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// EnvCmd ...
var EnvCmd = &cobra.Command{
	Use:    "env",
	Short:  "Prints the environment variables.",
	Hidden: false,
	Run: func(cmd *cobra.Command, args []string) {
		for _, env := range os.Environ() {
			fmt.Println(env)
		}
	},
}
