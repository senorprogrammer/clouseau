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

	railsConfig := modules.NewRailsConfigScanner(path)
	railsConfig.Load(path)

	// display := display.NewTableData(railsConfig)
	// display.Render()

	display := display.NewHtmlData(railsConfig)
	display.Render()
	display.Show()
}
