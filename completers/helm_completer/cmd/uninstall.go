package cmd

import (
	"time"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/helm"
	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:     "uninstall",
	Short:   "uninstall a release",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(uninstallCmd).Standalone()
	uninstallCmd.Flags().String("description", "", "add a custom description")
	uninstallCmd.Flags().Bool("dry-run", false, "simulate a uninstall")
	uninstallCmd.Flags().Bool("keep-history", false, "remove all associated resources and mark the release as deleted, but retain the release history")
	uninstallCmd.Flags().Bool("no-hooks", false, "prevent hooks from running during uninstallation")
	uninstallCmd.Flags().Duration("timeout", 5*time.Minute, "time to wait for any individual Kubernetes operation (like Jobs for hooks)")
	rootCmd.AddCommand(uninstallCmd)

	carapace.Gen(uninstallCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return helm.ActionReleases(helm.ReleasesOpts{
				Namespace:   rootCmd.Flag("namespace").Value.String(),
				KubeContext: rootCmd.Flag("kube-context").Value.String(),
			})
		}),
	)
}
