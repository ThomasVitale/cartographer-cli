package cmd

import (
	"github.com/ThomasVitale/cartographer-cli/cmd/apps"
	"github.com/ThomasVitale/cartographer-cli/cmd/cluster"
	"github.com/ThomasVitale/cartographer-cli/cmd/test"

	"github.com/spf13/cobra"

	"os"
)

var rootCmd = &cobra.Command{
	Use:   "carto",
	Short: "Manage Cartographer installation, resources and testing.",
	Long:  `Provides features to manage a Cartographer installation, work with its CRDs and test them.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(apps.AppsCmd)
	rootCmd.AddCommand(cluster.ClusterCmd)
	rootCmd.AddCommand(test.TestCmd)
}
