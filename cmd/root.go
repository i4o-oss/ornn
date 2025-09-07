package cmd

import (
	"context"
	"os"

	"github.com/charmbracelet/fang"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ornn",
	Short: "Batteries-included project generator for Go",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := fang.Execute(context.Background(), rootCmd, fang.WithoutCompletions(), fang.WithoutManpage()); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
		Run:    func(cmd *cobra.Command, args []string) {}, // no-op run function
	})
	rootCmd.AddCommand(CreateCmd)
}
