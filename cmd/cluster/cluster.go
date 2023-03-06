package cluster

import (
	"github.com/spf13/cobra"
)

var ClusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Manage a Cartographer installation on a Kubernetes cluster",
	Long:  `Provides features to install and manage a Cartographer installation on a Kubernetes cluster.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	// Here you will define your flags and configuration settings.
}
