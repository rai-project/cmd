package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// licenseCmd represents the license command
var LicenseCmd = &cobra.Command{
	Use:   "license",
	Short: "Display the project license.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(_escFSMustString(false, "/LICENSE.TXT"))
	},
}
