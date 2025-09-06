/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/i4o-oss/ornn/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new project",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return config.Init(cmd)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		router := viper.GetString("router")
		fmt.Printf("create command called, using router: %s\n", router)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.PersistentFlags().StringVar(&config.ConfigurationFile, "config", "", "config file (default is $HOME/.ornn.toml)")

	createCmd.Flags().String("router", "net/http", "Router to use for the project")
}
