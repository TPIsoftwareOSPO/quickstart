package cmd

import (
	"fmt"
	"github.com/TPIsoftwareOSPO/quickstart/app"
	"github.com/TPIsoftwareOSPO/quickstart/config"
	"github.com/TPIsoftwareOSPO/quickstart/utils"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Generate minimal quickstart.yaml file",
	Long:  "Generate minimal quickstart.yaml in the current directory with a basic 'echo' task.",
	Run: func(cmd *cobra.Command, args []string) {

		if dir, err := os.Getwd(); err == nil {
			message := fmt.Sprintf("dir: %s", dir)
			utils.SharedAppLogger.Info(message)
		}

		var echoConfig map[string]any

		if app.InitCmdIsWindows {
			echoConfig = map[string]any{
				"tasks": []map[string]any{
					{
						"name":       "echo",
						"executable": "cmd.exe",
						"args":       []string{"/C", "echo hello world"},
					},
				},
			}
		} else {
			echoConfig = map[string]any{
				"tasks": []map[string]any{
					{
						"name":       "echo",
						"executable": "echo",
						"args":       []string{"hello", "world"},
					},
				},
			}
		}

		currentDir, err := os.Getwd()
		if err != nil {
			utils.SharedAppLogger.Fatal(err)
		}

		var outputFile = config.GetDefaultFileNameWithExtension()

		if app.InitCmdOutput != "" {
			outputFile = app.InitCmdOutput
		}

		outputFile = filepath.Join(currentDir, outputFile)

		if _, err := os.Stat(outputFile); err == nil {
			utils.SharedAppLogger.Warn(fmt.Sprintf("'quickstart.yaml' already exists in '%s'. Aborting to prevent overwrite.\n", currentDir))
			utils.SharedAppLogger.Warn("If you wish to overwrite, please delete the existing file first.")
			return
		} else if !os.IsNotExist(err) {
			// Handle other potential errors from os.Stat
			utils.SharedAppLogger.Fatal(err)
		}

		yamlContent, err := yaml.Marshal(echoConfig)
		if err != nil {
			utils.SharedAppLogger.Fatal(err)
		}

		err = os.WriteFile(outputFile, yamlContent, 0644)
		if err != nil {
			utils.SharedAppLogger.Fatal(err)
		}
		utils.SharedAppLogger.Info("Successfully generated quickstart.yaml", "path", outputFile)
	},
}

func init() {
	InitCmd.PersistentFlags().StringVarP(&app.InitCmdOutput, "output", "o", "", "Specify the generated configuration file name")
	InitCmd.PersistentFlags().BoolVar(&app.InitCmdIsWindows, "win", false, "Use windows configuration")
}
