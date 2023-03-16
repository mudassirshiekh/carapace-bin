package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var autoscaleCmd = &cobra.Command{
	Use:     "autoscale",
	Short:   "Auto-scale a deployment, replica set, stateful set, or replication controller",
	GroupID: "deploy",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaleCmd).Standalone()
	autoscaleCmd.Flags().Bool("allow-missing-template-keys", true, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	autoscaleCmd.Flags().Int32("cpu-percent", -1, "The target average CPU utilization (represented as a percent of requested CPU) over all the pods. If it's not specified or negative, a default autoscaling policy will be used.")
	autoscaleCmd.Flags().String("dry-run", "none", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.")
	autoscaleCmd.Flags().String("field-manager", "kubectl-autoscale", "Name of the manager used to track field ownership.")
	autoscaleCmd.Flags().StringSliceP("filename", "f", []string{}, "Filename, directory, or URL to files identifying the resource to autoscale.")
	autoscaleCmd.Flags().StringP("kustomize", "k", "", "Process the kustomization directory. This flag can't be used together with -f or -R.")
	autoscaleCmd.Flags().Int32("max", -1, "The upper limit for the number of pods that can be set by the autoscaler. Required.")
	autoscaleCmd.Flags().Int32("min", -1, "The lower limit for the number of pods that can be set by the autoscaler. If it's not specified or negative, the server will apply a default value.")
	autoscaleCmd.Flags().String("name", "", "The name for the newly created object. If not specified, the name of the input resource will be used.")
	autoscaleCmd.Flags().StringP("output", "o", "", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).")
	autoscaleCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.")
	autoscaleCmd.Flags().Bool("save-config", false, "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.")
	autoscaleCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	autoscaleCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].")
	autoscaleCmd.Flag("dry-run").NoOptDefVal = "unchanged"
	rootCmd.AddCommand(autoscaleCmd)

	// TODO field-manager
	carapace.Gen(autoscaleCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":   action.ActionDryRunModes(),
		"filename":  carapace.ActionFiles(),
		"kustomize": carapace.ActionDirectories(),
		"output":    action.ActionOutputFormats(),
		"template":  carapace.ActionFiles(),
	})

	carapace.Gen(autoscaleCmd).PositionalCompletion(
		carapace.ActionValues("deployments", "replicasets", "replicationcontrollers"),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionResources("", c.Args[0])
		}),
	)
}
