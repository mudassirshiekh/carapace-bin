package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "print Go environment information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(envCmd).Standalone()

	envCmd.Flags().Bool("json", false, "print the environment in JSON format")
	envCmd.Flags().BoolS("u", "u", false, "unsets the default setting for the named environment variables")
	envCmd.Flags().StringS("w", "w", "", "changes the default settings of the named environment variables")
	rootCmd.AddCommand(envCmd)
}
