package main

import (
	"flag"
	"fmt"

	"github.com/senorprogrammer/conf_check/display"
	"github.com/senorprogrammer/conf_check/modules"
)

func main() {
	path := flag.String("dir", "./", "Path to Rails application")
	flag.Parse()

	fmt.Println("Running conf_check...")

	/* RailsConfig checking */
	railsConfChecker := modules.NewRailsConfigChecker(path)
	railsConfChecker.Run()

	/* Config checking */
	configChecker := modules.NewConfigChecker("Config", *path, `AppConfig\.?[A-Za-z_]+`)
	configChecker.Run()

	/* Figaro checking */
	figaroChecker := modules.NewConfigChecker("Figaro", *path, `Figaro\.env?[A-Za-z._]+`)
	figaroChecker.Run()

	/* Environment variable checking */
	envVarChecker := modules.NewConfigChecker("ENV Vars", *path, `ENV\[(.*?)\]`)
	envVarChecker.Run()

	/* HTML rendering */
	display := display.NewHtmlData(envVarChecker, configChecker, figaroChecker, railsConfChecker)
	display.Render()
	display.Show()
}
