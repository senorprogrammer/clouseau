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
	railsConfChecker.Load()

	/* Environment variable checking */
	envVarChecker := modules.NewEnvVarChecker(path)
	envVarChecker.Load()

	/* HTML rendering */
	display := display.NewHtmlData(envVarChecker, railsConfChecker)
	display.Render()
	display.Show()
}
