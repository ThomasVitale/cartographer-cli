package cluster

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version string
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a cluster with Cartographer",
	Long:  `Install a specific version of Cartographer and its dependencies on a Kubernetes cluster.`,
	Run: func(cmd *cobra.Command, args []string) {
		installCartographer()
	},
}

func init() {
	initCmd.Flags().StringVarP(&version, "version", "v", "latest", "The version of Cartographer to install")
	ClusterCmd.AddCommand(initCmd)
}

func installCartographer() {
	// We could install Cartographer. Similar approach as https://github.com/carvel-dev/kapp-controller/issues/1043
	fmt.Println("This functionality has not been implemented yet.")
}
