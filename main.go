package main

import (
	"flag"
	"fmt"

	"github.com/senorprogrammer/conf_check/display"
	"github.com/senorprogrammer/conf_check/modules"
)

/* -------------------- Main -------------------- */

func main() {
	path := flag.String("dir", "./", "Path to Rails application")
	flag.Parse()

	fmt.Println("Running conf_check...")

	railsConfChecker := modules.NewRailsConfigChecker(path)
	configChecker := modules.NewConfigChecker("Config", *path, `AppConfig\.?[A-Za-z_]+`)
	envVarChecker := modules.NewConfigChecker("ENV Vars", *path, `ENV\[(.*?)\]`)
	figaroChecker := modules.NewConfigChecker("Figaro", *path, `Figaro\.env?[A-Za-z._]+`)

	checkbox := modules.Checkbox{Path: *path, RailsConfigCheck: railsConfChecker}
	checkbox.Append(configChecker)
	checkbox.Append(figaroChecker)
	checkbox.Append(envVarChecker)
	checkbox.Run()

	/* HTML rendering */
	display := display.NewHtmlData(envVarChecker, configChecker, figaroChecker, railsConfChecker)
	display.Render()
	display.Show()
}
