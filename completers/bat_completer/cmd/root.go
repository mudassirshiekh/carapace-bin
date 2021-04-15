package cmd

import (
	"os/exec"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bat",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Root()
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("color", "", "Specify when to use colored output.")
	rootCmd.Flags().String("decorations", "", "Specify when to use the decorations that have been specified via '--style'.")
	rootCmd.Flags().BoolP("diff", "d", false, "Only show lines that have been added/removed/modified with respect to the Git index.")
	rootCmd.Flags().String("diff-context", "", "Include N lines of context around added/removed/modified lines when using '--diff'.")
	rootCmd.Flags().String("file-name", "", "Specify the name to display for a file.")
	rootCmd.Flags().BoolP("help", "h", false, "Print this help message.")
	rootCmd.Flags().StringP("highlight-line", "H", "", "Highlight the specified line ranges with a different background color.")
	rootCmd.Flags().String("italic-text", "", "Specify when to use ANSI sequences for italic text in the output.")
	rootCmd.Flags().StringP("language", "l", "", "Explicitly set the language for syntax highlighting.")
	rootCmd.Flags().StringP("line-range", "r", "", "Only print the specified range of lines for each file.")
	rootCmd.Flags().BoolP("list-languages", "L", false, "Display a list of supported languages for syntax highlighting.")
	rootCmd.Flags().Bool("list-themes", false, "Display a list of supported themes for syntax highlighting.")
	rootCmd.Flags().StringP("map-syntax", "m", "", "Map a glob pattern to an existing syntax name.")
	rootCmd.Flags().BoolP("number", "n", false, "Only show line numbers, no other decorations.")
	rootCmd.Flags().String("pager", "", "Determine which pager is used.")
	rootCmd.Flags().String("paging", "", "Specify when to use the pager.")
	rootCmd.Flags().BoolP("plain", "p", false, "Only show plain style, no decorations. This is an alias for '--style=plain'.")
	rootCmd.Flags().BoolP("show-all", "A", false, "Show non-printable characters like space, tab or newline.")
	rootCmd.Flags().String("style", "", "Configure which elements to display in addition to the file contents.")
	rootCmd.Flags().String("tabs", "", "Set the tab width to T spaces.")
	rootCmd.Flags().String("terminal-width", "", "Explicitly set the width of the terminal instead of determining it automatically.")
	rootCmd.Flags().String("theme", "", "Set the theme for syntax highlighting.")
	rootCmd.Flags().BoolP("unbuffered", "u", false, "This option exists for POSIX-compliance reasons ('u' is for 'unbuffered').")
	rootCmd.Flags().BoolP("version", "V", false, "Show version information.")
	rootCmd.Flags().String("wrap", "", "Specify the text-wrapping mode (*auto*, never, character).")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"color":       carapace.ActionValues("auto", "never", "always"),
		"decorations": carapace.ActionValues("auto", "never", "always"),
		"italic-text": carapace.ActionValues("never", "always"),
		"language":    ActionLanguage(),
		"style":       carapace.ActionValues("auto", "full", "plain", "changes", "header", "grid", "numbers", "snip"),
		"theme":       ActionTheme(),
		"wrap":        carapace.ActionValues("auto", "never", "character"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}

func ActionLanguage() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if output, err := exec.Command("bat", "--list-languages").Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			values := make([]string, 0)
			for _, line := range strings.Split(string(output), "\n") {
				// TODO fix ':' character in carapace
				splitted := strings.Split(line, ":")
				if len(splitted) == 1 {
					values = append(values, splitted[0], "")
				} else if len(splitted) > 1 {
					values = append(values, splitted[0], splitted[1])
				}
			}
			return carapace.ActionValuesDescribed(values...)
		}
	})
}

func ActionTheme() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if output, err := exec.Command("bat", "--list-themes").Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			return carapace.ActionValues(strings.Split(string(output), "\n")...)
		}
	})
}
