package cluster

import (
	"fmt"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Gathers information about a Cartographer installation",
	Long:  `Gathers information about a Cartographer installation on a Kubernetes cluster.`,
	Run: func(cmd *cobra.Command, args []string) {
		getCartographerVersion()
	},
}

func init() {
	ClusterCmd.AddCommand(infoCmd)
}

func getCartographerVersion() {
	// Read from `kubectl get deploy/cartographer-controller -n cartographer-system -o yaml`, the label is `app.kubernetes.io/version`
	fmt.Println("This functionality has not been implemented yet.")
}
