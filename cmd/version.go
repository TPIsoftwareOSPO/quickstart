package cmd

import (
	"fmt"
	"github.com/TPIsoftwareOSPO/quickstart/app"
	"github.com/TPIsoftwareOSPO/quickstart/utils"
	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version number and build details of quickstart",
	Run: func(cmd *cobra.Command, args []string) {

		var info = fmt.Sprintf("[%s] (Build: %s Commit: %s Portable: %s)",
			app.Version,
			app.BuildDate,
			app.CommitHash,
			app.Portable)
		utils.SharedAppLogger.Info(info)
	},
}
