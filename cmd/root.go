package cmd

import (
	"github.com/TPIsoftwareOSPO/quickstart/app"
	"github.com/TPIsoftwareOSPO/quickstart/utils"
	"github.com/spf13/cobra"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "dgrkt",
	Short: "A versatile application launcher with flexible configuration options",
	Long: `dgrkt is a versatile command-line tool designed to simplify the management of complex, multi-process applications or task sequences.
	
It allows you to define a series of commands (tasks) and their operational parameters within a single, human-readable YAML configuration file.
This includes specifying the executable, arguments, working directories, environment variables, and crucial inter-task dependencies.

A core feature of dgrkt is its robust health check mechanism. Each task can be configured with HTTP-based health checks (including JSON response validation via JSONPath) or command-line script-based health checks.
Tasks with dependencies will only start once their prerequisites are deemed healthy, ensuring a stable and reliable startup order for your services.

Whether you're spinning up a local development environment with multiple microservices, orchestrating integration tests, or automating complex workflows,
dgrkt provides a declarative and efficient way to manage your system's components.
	`,
}

func Execute() {
	RootCmd.AddCommand(CheckCmd)
	RootCmd.AddCommand(UpCmd)
	RootCmd.AddCommand(DownCmd)
	RootCmd.AddCommand(VersionCmd)
	RootCmd.AddCommand(InitCmd)
	if len(os.Args) == 1 && app.Portable == "true" {
		RootCmd.SetArgs([]string{UpCmd.Use})
	}
	if err := RootCmd.Execute(); err != nil {
		utils.SharedAppLogger.Fatal(err)
	}
}

func init() {
	if app.Portable == "true" {
		cobra.MousetrapHelpText = ""
	}
	RootCmd.PersistentFlags().BoolVar(&app.DebugMode, "debug", false, "Enabling debug mode will display more detailed console logs.")
}
