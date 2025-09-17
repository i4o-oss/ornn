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
			fmt.Printf("create command called\n\n")
			fmt.Println("project settings")
			fmt.Printf("name: %s\n", viper.GetString("name"))
			fmt.Printf("path: %s\n\n", viper.GetString("path"))

			fmt.Println("project configuration")
			fmt.Printf("auth: %s\n", viper.GetString("auth"))
			fmt.Printf("cache: %s\n", viper.GetString("cache"))
			fmt.Printf("database: %s\n", viper.GetString("database"))
			fmt.Printf("file-storage: %s\n", viper.GetString("file-storage"))
			fmt.Printf("queue: %s\n", viper.GetString("queue"))
			fmt.Printf("rate-limiting: %s\n", viper.GetString("rate-limiting"))
			fmt.Printf("router: %s\n", viper.GetString("router"))
			return nil
		},
	}
	projectSettings = config.ProjectSettings{
		Configuration: configuration,
	}
)

func init() {
	// Hide the help flag and call the parent help command
	CreateCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		cmd.Flags().MarkHidden("help")
		cmd.Parent().HelpFunc()(cmd, args)
	})

	// Project Settings
	CreateCmd.Flags().StringVar(&projectSettings.Name, "name", "", "Name of the project")
	CreateCmd.Flags().StringVar(&projectSettings.Path, "path", ".", "Path of the project")

	/*
		// File-based configuration
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		CreateCmd.Flags().StringVar(&config.ConfigurationFile, "config", home+"/.ornn.toml", "config file location")
	*/

	// Flag-based configuration
	CreateCmd.Flags().StringVar(&projectSettings.Configuration.Auth, "auth", "Email/Password", "Email/Password / Magic Link")
	CreateCmd.Flags().StringVar(&projectSettings.Configuration.Cache, "cache", "", "Redis / Memcached / Valkey")
	CreateCmd.Flags().StringVar(&projectSettings.Configuration.Database, "database", "PostgreSQL", "PostgreSQL / SQLite")
	CreateCmd.Flags().StringVar(&projectSettings.Configuration.FileStorage, "file-storage", "", "Cloudflare R2 / AWS S3")
	CreateCmd.Flags().StringVar(&projectSettings.Configuration.Queue, "queue", "", "Redis / AWS SQS")
	CreateCmd.Flags().StringVar(&projectSettings.Configuration.RateLimiting, "rate-limiting", "In-memory", "In-memory / Redis / Memcached / Valkey")
	CreateCmd.Flags().StringVar(&projectSettings.Configuration.Router, "router", "net/http", "BunRouter / Chi / Echo / Fiber / Gin / Gorilla Mux / net/http")
	CreateCmd.Flags().BoolVar(&projectSettings.Configuration.Tasks, "tasks", false, "Set up Tasks")
	CreateCmd.Flags().BoolVar(&projectSettings.Configuration.Scheduler, "scheduler", false, "Set up Scheduler")
	CreateCmd.Flags().BoolVar(&projectSettings.Configuration.SMTP, "smtp", false, "Set up SMTP")
	CreateCmd.Flags().BoolVar(&projectSettings.Configuration.Webhooks, "webhooks", false, "Set up Webhooks")
}
