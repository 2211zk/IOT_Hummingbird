package cmd

import (
	"cobra-script-center/internal/models"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var executionCmd = &cobra.Command{
	Use:   "execution",
	Short: "Manage script executions",
	Long:  `View and manage script execution history and status.`,
}

var executionListCmd = &cobra.Command{
	Use:   "list",
	Short: "List script executions",
	Long:  `List script executions with optional filtering.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		application, err := initApp()
		if err != nil {
			return fmt.Errorf("failed to initialize app: %w", err)
		}

		scriptName, _ := cmd.Flags().GetString("script")
		limit, _ := cmd.Flags().GetInt("limit")

		var executions []*models.Execution

		if scriptName != "" {
			// Get executions for specific script
			script, err := application.ScriptService.GetScriptByName(scriptName)
			if err != nil {
				return fmt.Errorf("script not found: %w", err)
			}
			executions, err = application.ExecutionService.GetExecutionsByScript(script.ID, limit)
		} else {
			// Get all running executions for now
			executions, err = application.ExecutionService.GetRunningExecutions()
		}

		if err != nil {
			return fmt.Errorf("failed to list executions: %w", err)
		}

		if len(executions) == 0 {
			fmt.Println("No executions found.")
			return nil
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "ID\tSCRIPT_ID\tSTATUS\tSTART_TIME\tDURATION\tUSER_ID")
		fmt.Fprintln(w, "---\t---------\t------\t----------\t--------\t-------")

		for _, execution := range executions {
			startTime := "N/A"
			if execution.StartTime != nil {
				startTime = execution.StartTime.Format("15:04:05")
			}

			duration := "N/A"
			if execution.StartTime != nil && execution.EndTime != nil {
				duration = execution.Duration().String()
			}

			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\n",
				execution.ID[:8],
				execution.ScriptID[:8],
				execution.Status,
				startTime,
				duration,
				execution.UserID[:8],
			)
		}

		return w.Flush()
	},
}

var executionShowCmd = &cobra.Command{
	Use:   "show [execution-id]",
	Short: "Show execution details",
	Long:  `Show detailed information about a specific execution.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		executionID := args[0]

		application, err := initApp()
		if err != nil {
			return fmt.Errorf("failed to initialize app: %w", err)
		}

		execution, err := application.ExecutionService.GetExecution(executionID)
		if err != nil {
			return fmt.Errorf("execution not found: %w", err)
		}

		fmt.Printf("Execution ID: %s\n", execution.ID)
		fmt.Printf("Script ID: %s\n", execution.ScriptID)
		fmt.Printf("User ID: %s\n", execution.UserID)
		fmt.Printf("Status: %s\n", execution.Status)
		fmt.Printf("Created: %s\n", execution.CreatedAt.Format("2006-01-02 15:04:05"))

		if execution.StartTime != nil {
			fmt.Printf("Started: %s\n", execution.StartTime.Format("2006-01-02 15:04:05"))
		}

		if execution.EndTime != nil {
			fmt.Printf("Ended: %s\n", execution.EndTime.Format("2006-01-02 15:04:05"))
			fmt.Printf("Duration: %s\n", execution.Duration().String())
		}

		if len(execution.Params) > 0 {
			fmt.Println("\nParameters:")
			for key, value := range execution.Params {
				fmt.Printf("  %s: %s\n", key, value)
			}
		}

		if execution.Output != "" {
			fmt.Printf("\nOutput:\n%s\n", execution.Output)
		}

		if execution.Error != "" {
			fmt.Printf("\nError:\n%s\n", execution.Error)
		}

		return nil
	},
}

var executionCancelCmd = &cobra.Command{
	Use:   "cancel [execution-id]",
	Short: "Cancel a running execution",
	Long:  `Cancel a running script execution.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		executionID := args[0]

		application, err := initApp()
		if err != nil {
			return fmt.Errorf("failed to initialize app: %w", err)
		}

		err = application.ExecutionService.CancelExecution(executionID)
		if err != nil {
			return fmt.Errorf("failed to cancel execution: %w", err)
		}

		fmt.Printf("Execution '%s' cancelled successfully\n", executionID)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(executionCmd)

	// Add subcommands
	executionCmd.AddCommand(executionListCmd)
	executionCmd.AddCommand(executionShowCmd)
	executionCmd.AddCommand(executionCancelCmd)

	// List command flags
	executionListCmd.Flags().StringP("script", "s", "", "Filter by script name")
	executionListCmd.Flags().IntP("limit", "l", 10, "Limit number of results")
}
