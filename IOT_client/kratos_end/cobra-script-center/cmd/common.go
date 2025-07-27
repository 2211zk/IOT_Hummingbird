package cmd

import "cobra-script-center/internal/app"

// initApp initializes the application
func initApp() (*app.App, error) {
	return app.NewApp()
}
