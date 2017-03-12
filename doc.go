package cmd

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/Unknwon/com"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	jww "github.com/spf13/jwalterweatherman"
)

const gendocFrontmatterTemplate = `---
date: %s
title: "%s"
slug: %s
url: %s
---
`

var (
	gendocdir string
)

var GendocCmd = &cobra.Command{
	Use:   "doc",
	Short: "Generate Markdown documentation for the RAI CLI.",
	Long: `Generate Markdown documentation for the RAI CLI.
This command is, mostly, used to create up-to-date documentation
of RAI's command-line interface.
It creates one Markdown file per command with front matter suitable
for rendering in Hugo.`,

	RunE: func(cmd *cobra.Command, args []string) error {
		gendocdir = filepath.Clean(gendocdir)
		if !strings.HasSuffix(gendocdir, string(os.PathSeparator)) {
			gendocdir += string(os.PathSeparator)
		}
		if !com.IsDir(gendocdir) {
			jww.FEEDBACK.Println("Directory", gendocdir, "does not exist, creating...")
			os.MkdirAll(gendocdir, os.ModePerm)
		}
		now := time.Now().Format(time.RFC3339)
		prepender := func(filename string) string {
			name := filepath.Base(filename)
			base := strings.TrimSuffix(name, path.Ext(name))
			url := "/commands/" + strings.ToLower(base) + "/"
			return fmt.Sprintf(gendocFrontmatterTemplate, now, strings.Replace(base, "_", " ", -1), base, url)
		}

		linkHandler := func(name string) string {
			base := strings.TrimSuffix(name, path.Ext(name))
			return "/commands/" + strings.ToLower(base) + "/"
		}

		jww.FEEDBACK.Println("Generating RAI command-line documentation in", gendocdir, "...")
		doc.GenMarkdownTreeCustom(cmd.Root(), gendocdir, prepender, linkHandler)
		jww.FEEDBACK.Println("Done.")

		return nil
	},
}

func init() {
	GendocCmd.PersistentFlags().StringVarP(&gendocdir, "dir", "o", "/tmp/raidoc/", "the directory to write the doc.")

	// For bash-completion
	GendocCmd.PersistentFlags().SetAnnotation("dir", cobra.BashCompSubdirsInDir, []string{})
}
