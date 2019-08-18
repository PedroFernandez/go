package main

import (
	"github.com/PedroFernandez/go/cli"
	"github.com/spf13/cobra"
)
func main() {
	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(cli.InitBeersCmd())
	rootCmd.AddCommand(cli.InitStoresCmd())
	rootCmd.Execute()
}
