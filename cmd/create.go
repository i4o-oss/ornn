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

var (
	configuration config.Configuration
	CreateCmd     = &cobra.Command{
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
)

func init() {
	// Hide the help flag and call the parent help command
	CreateCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		cmd.Flags().MarkHidden("help")
		cmd.Parent().HelpFunc()(cmd, args)
	})

	// File-based configuration
	CreateCmd.Flags().StringVar(&config.ConfigurationFile, "config", "", "config file (default is $HOME/.ornn.toml)")

	// Flag-based configuration
	CreateCmd.Flags().StringVar(&configuration.Cache, "cache", "Redis", "Redis / Memcached / Valkey")
	CreateCmd.Flags().StringVar(&configuration.Database, "database", "PostgreSQL", "PostgreSQL / SQLite")
	CreateCmd.Flags().StringVar(&configuration.FileStorage, "file-storage", "AWS S3", "Cloudflare R2 / AWS S3")
	CreateCmd.Flags().StringVar(&configuration.Queue, "queue", "Redis", "AWS SQS / Redis")
	CreateCmd.Flags().StringVar(&configuration.RateLimiting, "rate-limiting", "Redis", "Redis / Memcached / Valkey")
	CreateCmd.Flags().StringVar(&configuration.Router, "router", "net/http", "BunRouter / Chi / Echo / Fiber / Gin / Gorilla Mux / net/http")
	CreateCmd.Flags().BoolVar(&configuration.SMTP, "smtp", false, "Set up SMTP")
	CreateCmd.Flags().BoolVar(&configuration.Webhooks, "webhooks", false, "Set up Webhooks")
}
