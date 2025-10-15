/*
Copyright © 2025 Vulcan Shen vulcan.shen@tpisoftware.com
*/
package main

import (
	"github.com/TPIsoftwareOSPO/quickstart/app"
	"github.com/TPIsoftwareOSPO/quickstart/cmd"
	"github.com/TPIsoftwareOSPO/quickstart/utils"
	"os"
	"path/filepath"
)

func main() {

	if len(os.Args) == 1 && app.Portable == "true" {
		utils.SharedAppLogger.Info("portable mode")
		// portable mode
		path, err := os.Executable()
		if err != nil {
			utils.SharedAppLogger.Fatal(err)
		}
		if err = os.Chdir(filepath.Dir(path)); err != nil {
			utils.SharedAppLogger.Fatal(err)
		}

	} else {
		// cli mode
		dir, err := os.Getwd()
		if err != nil {
			utils.SharedAppLogger.Fatal(err)
		}
		if err = os.Chdir(dir); err != nil {
			utils.SharedAppLogger.Fatal(err)
		}
	}

	cmd.Execute()
}
