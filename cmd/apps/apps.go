package apps

import (
	"github.com/spf13/cobra"

	"context"
	"fmt"
	"strings"

	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	_ "k8s.io/client-go/plugin/pkg/client/auth/openstack"

	"github.com/fatih/color"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"

	cartov1alpha1 "github.com/vmware-tanzu/apps-cli-plugin/pkg/apis/cartographer/v1alpha1"
	knativeservingv1 "github.com/vmware-tanzu/apps-cli-plugin/pkg/apis/knative/serving/v1"
	tanzucli "github.com/vmware-tanzu/apps-cli-plugin/pkg/cli-runtime"
	"github.com/vmware-tanzu/apps-cli-plugin/pkg/cli-runtime/logs"
	"github.com/vmware-tanzu/apps-cli-plugin/pkg/cli-runtime/printer"
	"github.com/vmware-tanzu/apps-cli-plugin/pkg/commands"
	"github.com/vmware-tanzu/apps-cli-plugin/pkg/flags"
	"github.com/vmware-tanzu/apps-cli-plugin/pkg/logger"
)

var AppsCmd = &cobra.Command{
	Use:   "apps",
	Short: "Manage workloads and supply chains on Kubernetes",
	Long:  `Manage Cartographer resources on Kubernetes as part of your path to production, including Workload, Deliverable, SupplyChain and Delivery.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var (
	scheme = runtime.NewScheme()
)

func init() {
	// Configure scheme with Kubernetes Resources and CRDs used by Cartographer.
	_ = clientgoscheme.AddToScheme(scheme)
	_ = cartov1alpha1.AddToScheme(scheme)
	_ = knativeservingv1.AddToScheme(scheme)

	// Set up context for subcommands.
	context := context.Background()
	AppsCmd.CompletionOptions.DisableDefaultCmd = true
	context = logs.StashTailer(context, &logs.SternTailer{})
	cliConfig := tanzucli.Initialize(fmt.Sprintf("carto %s", AppsCmd.Use), scheme)

	// Add subcommands from Tanzu Apps CLI.
	AppsCmd.AddCommand(commands.NewClusterSupplyChainCommand(context, cliConfig))
	AppsCmd.AddCommand(commands.NewWorkloadCommand(context, cliConfig))
	AppsCmd.AddCommand(commands.NewDocsCommand(context, cliConfig))

	// Configure persistent flags which will work for this command and all subcommands.
	AppsCmd.PersistentFlags().StringVar(&cliConfig.KubeConfigFile, tanzucli.StripDash(flags.KubeConfigFlagName), "", "kubeconfig `file` (default is $HOME/.kube/config)")
	AppsCmd.MarkFlagFilename(tanzucli.StripDash(flags.KubeConfigFlagName))
	AppsCmd.PersistentFlags().StringVar(&cliConfig.CurrentContext, tanzucli.StripDash(flags.ContextFlagName), "", "`name` of the kubeconfig context to use (default is current-context defined by kubeconfig)")
	AppsCmd.PersistentFlags().BoolVar(&color.NoColor, tanzucli.StripDash(flags.NoColorFlagName), color.NoColor, "deactivate color, bold, animations, and emoji output")
	AppsCmd.PersistentFlags().Int32VarP(cliConfig.Verbose, tanzucli.StripDash(flags.VerboseLevelFlagName), "v", 1, "number for the log level verbosity")

	cobra.OnInitialize(func() {
		// Sync config to deactivate emojis printing
		cliConfig.NoColor = color.NoColor
		// Set the default logger
		cliConfig.SetLogger(logger.NewSinkLogger(cliConfig.Name, cliConfig.Verbose, cliConfig.Stderr))
	})

	// Override usage template to add arguments.
	AppsCmd.SetUsageTemplate(strings.ReplaceAll(AppsCmd.UsageTemplate(), "{{.UseLine}}", "{{useLine .}}"))
	cobra.AddTemplateFunc("useLine", func(cmd *cobra.Command) string {
		result := cmd.UseLine()
		flags := ""
		if strings.HasSuffix(result, " [flags]") {
			flags = " [flags]"
			result = result[0 : len(result)-len(flags)]
		}
		return result + tanzucli.FormatArgs(cmd) + flags
	})

	// Override default colors.
	printer.InfoColor = color.New(color.FgCyan, color.Bold)
	printer.SuccessColor = color.New(color.FgGreen, color.Bold)
	printer.WarnColor = color.New(color.FgYellow, color.Bold)
	printer.ErrorColor = color.New(color.FgRed, color.Bold)

	// AppsCmd.SilenceErrors = true
	// if err := AppsCmd.Execute(); err != nil {
	// 	// silent errors should not log, but still exit with an error code
	// 	// typically the command has already been logged with more detail
	// 	if !errors.Is(err, cli.SilentError) {
	// 		if aggregate, ok := err.(utilerrors.Aggregate); ok {
	// 			for _, err := range aggregate.Errors() {
	// 				cliConfig.Eprintf("%s %s\n", printer.Serrorf("Error:"), err)
	// 			}
	// 		} else if apierrors.IsForbidden(err) {
	// 			cliConfig.Eprintf("%s %s: %s\n", printer.Serrorf("Error:"), "Unable to complete command, you do not have permissions for this resource", err)
	// 		} else {
	// 			cliConfig.Eprintf("%s %s\n", printer.Serrorf("Error:"), err)
	// 		}
	// 	}
	// 	os.Exit(1)
	// }
}
