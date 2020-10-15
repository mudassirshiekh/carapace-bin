package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var mod_downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "download modules to local cache",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mod_downloadCmd).Standalone()

	mod_downloadCmd.Flags().BoolS("x", "x", false, "print the commands download executes")
	mod_downloadCmd.Flags().Bool("json", false, "print a sequence of JSON objects")
	modCmd.AddCommand(mod_downloadCmd)
}
