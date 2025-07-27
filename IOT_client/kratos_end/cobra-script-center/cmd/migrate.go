package cmd

import (
	"cobra-script-center/internal/database"
	"fmt"

	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
	Long:  `Run database migrations to set up or update the database schema.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Running database migrations...")

		if err := database.RunSimpleMigrations(); err != nil {
			return fmt.Errorf("failed to run migrations: %w", err)
		}

		fmt.Println("Database migrations completed successfully.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
