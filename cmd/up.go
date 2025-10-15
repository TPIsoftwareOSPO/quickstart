package cmd

import (
	"github.com/TPIsoftwareOSPO/quickstart/app"
	"github.com/TPIsoftwareOSPO/quickstart/config"
	"github.com/TPIsoftwareOSPO/quickstart/procedure"
	"github.com/TPIsoftwareOSPO/quickstart/utils"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
)

var AppTasks map[string]*procedure.Task
var UpCmd = &cobra.Command{
	Use:   "up",
	Short: "Execute tasks according to the YAML configuration file.",
	Long:  "Execute tasks according to the YAML configuration file. with command: quickstart up",
	PreRun: func(cmd *cobra.Command, args []string) {
		procedure.InitializeSpinnerAgent()
		procedure.StartSpinnerAgent()
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		procedure.StopSpinnerAgent()
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := CheckConfig(); err != nil {
			utils.SharedAppLogger.Fatal(err)
		}

		AppTasks = make(map[string]*procedure.Task)

		for _, taskConfig := range config.AppTasksConfig {
			var task, err = procedure.CreateTask(taskConfig)
			if err != nil {
				utils.SharedAppLogger.Fatal(err)
			}
			AppTasks[taskConfig.Name] = task
		}
		for _, taskConfig := range config.AppTasksConfig {
			var task = AppTasks[taskConfig.Name]
			if len(taskConfig.DependsOn) > 0 {
				for _, dependency := range taskConfig.DependsOn {
					task.AppendDependencies(AppTasks[dependency])
				}
			}
		}

		var waitGroup = &sync.WaitGroup{}
		waitGroup.Add(len(AppTasks))

		for _, task := range AppTasks {
			go task.Start(waitGroup)
		}

		waitGroup.Wait()

		if len(os.Args) == 1 && app.Portable == "true" && runtime.GOOS == "windows" {
			utils.SharedAppLogger.Info("Program completed, Press ctrl-c to exit.")
			c := make(chan os.Signal, 1)
			signal.Notify(c, os.Interrupt, syscall.SIGTERM)
			<-c
		}
	},
}

func init() {
	UpCmd.PersistentFlags().BoolVarP(&app.DetachMode, "detach", "d", false, "Launch tasks in the background")
	UpCmd.PersistentFlags().StringVarP(&app.TasksComposeFile, "configfile", "f", "", "Specify the path to the configuration file, default is 'quickstart.yaml'")
}
