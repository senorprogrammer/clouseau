package main

import (
	"flag"
	"fmt"

	"github.com/senorprogrammer/clouseau/display"
	"github.com/senorprogrammer/clouseau/modules"
)

/* -------------------- Main -------------------- */

func main() {
	path := flag.String("dir", "./", "Path to Rails application")
	flag.Parse()

	fmt.Println("Running clouseau...")

	railsConfChecker := modules.NewRailsConfigChecker(*path)
	railsConfChecker.Run()

	configChecker := modules.NewConfigChecker("Config", *path, `AppConfig\.?[A-Za-z_]+`)
	envVarChecker := modules.NewConfigChecker("ENV Vars", *path, `ENV\[(.*?)\]`)
	figaroChecker := modules.NewConfigChecker("Figaro", *path, `Figaro\.env?[A-Za-z._]+`)

	checkbox := modules.Checkbox{Path: *path}
	checkbox.Append(configChecker)
	checkbox.Append(figaroChecker)
	checkbox.Append(envVarChecker)
	checkbox.Run()

	/* HTML rendering */
	display := display.NewHtmlData(envVarChecker, configChecker, figaroChecker, railsConfChecker)
	display.Render()
	display.Show()
}
