package cmd

import (
	"cobra-script-center/internal/models"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var scriptCmd = &cobra.Command{
	Use:   "script",
	Short: "Manage scripts",
	Long:  `Create, list, edit, delete, and execute scripts.`,
}

var scriptCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new script",
	Long:  `Create a new script with specified name and language.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")
		language, _ := cmd.Flags().GetString("language")
		description, _ := cmd.Flags().GetString("description")
		tags, _ := cmd.Flags().GetStringSlice("tags")

		if name == "" {
			return fmt.Errorf("script name is required")
		}

		if !models.IsValidLanguage(language) {
			return fmt.Errorf("unsupported language: %s", language)
		}

		application, err := initApp()
		if err != nil {
			return fmt.Errorf("failed to initialize app: %w", err)
		}

		script := &models.Script{
			Name:        name,
			Description: description,
			Language:    language,
			Tags:        tags,
			Content:     getDefaultScriptContent(language),
		}

		createdScript, err := application.ScriptService.CreateScript(script, "system") // TODO: get current user
		if err != nil {
			return fmt.Errorf("failed to create script: %w", err)
		}

		fmt.Printf("Script '%s' created successfully with ID: %s\n", createdScript.Name, createdScript.ID)
		return nil
	},
}

var scriptListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all scripts",
	Long:  `List all scripts with their details.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		application, err := initApp()
		if err != nil {
			return fmt.Errorf("failed to initialize app: %w", err)
		}

		filter := &models.ScriptFilter{}
		if name, _ := cmd.Flags().GetString("name"); name != "" {
			filter.Name = name
		}
		if language, _ := cmd.Flags().GetString("language"); language != "" {
			filter.Language = language
		}
		if tags, _ := cmd.Flags().GetStringSlice("tags"); len(tags) > 0 {
			filter.Tags = tags
		}

		scripts, err := application.ScriptService.ListScripts(filter)
		if err != nil {
			return fmt.Errorf("failed to list scripts: %w", err)
		}

		if len(scripts) == 0 {
			fmt.Println("No scripts found.")
			return nil
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "ID\tNAME\tLANGUAGE\tTAGS\tCREATED\tACTIVE")
		fmt.Fprintln(w, "---\t----\t--------\t----\t-------\t------")

		for _, script := range scripts {
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%t\n",
				script.ID[:8],
				script.Name,
				script.Language,
				strings.Join(script.Tags, ","),
				script.CreatedAt.Format("2006-01-02"),
				script.IsActive,
			)
		}

		return w.Flush()
	},
}

var scriptRunCmd = &cobra.Command{
	Use:   "run [script-name]",
	Short: "Execute a script",
	Long:  `Execute a script by name or ID.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		scriptName := args[0]
		params, _ := cmd.Flags().GetStringToString("param")

		application, err := initApp()
		if err != nil {
			return fmt.Errorf("failed to initialize app: %w", err)
		}

		execution, err := application.ExecutionService.ExecuteScript(scriptName, "system", params) // TODO: get current user
		if err != nil {
			return fmt.Errorf("failed to execute script: %w", err)
		}

		fmt.Printf("Script execution started with ID: %s\n", execution.ID)
		fmt.Printf("Status: %s\n", execution.Status)

		if execution.Output != "" {
			fmt.Printf("Output:\n%s\n", execution.Output)
		}

		if execution.Error != "" {
			fmt.Printf("Error:\n%s\n", execution.Error)
		}

		return nil
	},
}

var scriptDeleteCmd = &cobra.Command{
	Use:   "delete [script-name]",
	Short: "Delete a script",
	Long:  `Delete a script by name or ID.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		scriptName := args[0]

		application, err := initApp()
		if err != nil {
			return fmt.Errorf("failed to initialize app: %w", err)
		}

		err = application.ScriptService.DeleteScript(scriptName)
		if err != nil {
			return fmt.Errorf("failed to delete script: %w", err)
		}

		fmt.Printf("Script '%s' deleted successfully\n", scriptName)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(scriptCmd)

	// Add subcommands
	scriptCmd.AddCommand(scriptCreateCmd)
	scriptCmd.AddCommand(scriptListCmd)
	scriptCmd.AddCommand(scriptRunCmd)
	scriptCmd.AddCommand(scriptDeleteCmd)

	// Create command flags
	scriptCreateCmd.Flags().StringP("name", "n", "", "Script name (required)")
	scriptCreateCmd.Flags().StringP("language", "l", "bash", "Script language")
	scriptCreateCmd.Flags().StringP("description", "d", "", "Script description")
	scriptCreateCmd.Flags().StringSliceP("tags", "t", []string{}, "Script tags")

	// List command flags
	scriptListCmd.Flags().StringP("name", "n", "", "Filter by name")
	scriptListCmd.Flags().StringP("language", "l", "", "Filter by language")
	scriptListCmd.Flags().StringSliceP("tags", "t", []string{}, "Filter by tags")

	// Run command flags
	scriptRunCmd.Flags().StringToStringP("param", "p", map[string]string{}, "Script parameters (key=value)")
}

func getDefaultScriptContent(language string) string {
	switch language {
	case "bash":
		return `#!/bin/bash
# Script description here
echo "Hello from bash script!"
`
	case "python":
		return `#!/usr/bin/env python3
# Script description here
print("Hello from Python script!")
`
	case "node":
		return `#!/usr/bin/env node
// Script description here
console.log("Hello from Node.js script!");
`
	case "go":
		return `package main

import "fmt"

func main() {
    fmt.Println("Hello from Go script!")
}
`
	case "powershell":
		return `# Script description here
Write-Host "Hello from PowerShell script!"
`
	default:
		return "# Add your script content here\n"
	}
}
