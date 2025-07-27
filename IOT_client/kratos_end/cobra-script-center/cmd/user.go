package cmd

import (
	"cobra-script-center/internal/models"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Manage users",
	Long:  `Create, list, and manage users.`,
}

var userCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new user",
	Long:  `Create a new user with specified username and role.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		username, _ := cmd.Flags().GetString("username")
		role, _ := cmd.Flags().GetString("role")
		password, _ := cmd.Flags().GetString("password")

		if username == "" {
			return fmt.Errorf("username is required")
		}

		if !models.IsValidRole(role) {
			return fmt.Errorf("invalid role: %s. Valid roles are: admin, user, viewer", role)
		}

		if password == "" {
			password = "changeme" // Default password
		}

		application, err := initApp()
		if err != nil {
			return fmt.Errorf("failed to initialize app: %w", err)
		}

		user := &models.User{
			Username: username,
			Role:     role,
		}

		createdUser, err := application.UserService.CreateUser(user, password)
		if err != nil {
			return fmt.Errorf("failed to create user: %w", err)
		}

		fmt.Printf("User '%s' created successfully with ID: %s\n", createdUser.Username, createdUser.ID)
		if password == "changeme" {
			fmt.Println("Default password 'changeme' assigned. Please change it on first login.")
		}
		return nil
	},
}

var userListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all users",
	Long:  `List all users with their details.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		application, err := initApp()
		if err != nil {
			return fmt.Errorf("failed to initialize app: %w", err)
		}

		users, err := application.UserService.ListUsers()
		if err != nil {
			return fmt.Errorf("failed to list users: %w", err)
		}

		if len(users) == 0 {
			fmt.Println("No users found.")
			return nil
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "ID\tUSERNAME\tROLE\tCREATED\tACTIVE")
		fmt.Fprintln(w, "---\t--------\t----\t-------\t------")

		for _, user := range users {
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%t\n",
				user.ID[:8],
				user.Username,
				user.Role,
				user.CreatedAt.Format("2006-01-02"),
				user.IsActive,
			)
		}

		return w.Flush()
	},
}

var userDeleteCmd = &cobra.Command{
	Use:   "delete [username]",
	Short: "Delete a user",
	Long:  `Delete a user by username.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		username := args[0]

		application, err := initApp()
		if err != nil {
			return fmt.Errorf("failed to initialize app: %w", err)
		}

		err = application.UserService.DeleteUser(username)
		if err != nil {
			return fmt.Errorf("failed to delete user: %w", err)
		}

		fmt.Printf("User '%s' deleted successfully\n", username)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(userCmd)

	// Add subcommands
	userCmd.AddCommand(userCreateCmd)
	userCmd.AddCommand(userListCmd)
	userCmd.AddCommand(userDeleteCmd)

	// Create command flags
	userCreateCmd.Flags().StringP("username", "u", "", "Username (required)")
	userCreateCmd.Flags().StringP("role", "r", "user", "User role (admin, user, viewer)")
	userCreateCmd.Flags().StringP("password", "p", "", "User password (default: changeme)")
}
