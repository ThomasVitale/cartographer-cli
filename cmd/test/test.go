package test

import (
	"fmt"

	"github.com/spf13/cobra"
)

var TestCmd = &cobra.Command{
	Use:   "test",
	Short: "Test Cartographer resources",
	Long:  `Provides features to test Cartographer Workloads and SupplyChains.`,
	Run: func(cmd *cobra.Command, args []string) {
		cartoTest()
	},
}

func init() {
	// Here you will define your flags and configuration settings.
}

func cartoTest() {
	// Hook the cartotest CLI here: https://github.com/vmware-tanzu/cartographer/blob/main/pkg/testing/cli.go
	fmt.Println("This functionality has not been implemented yet.")
}
