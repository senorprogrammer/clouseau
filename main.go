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

	railsConfig := modules.NewRailsConfig(path)
	railsConfig.Load(path)

	// table := display.Table{}
	// table.Render(railsConfig)

	display := display.NewHtmlData(railsConfig)
	display.Render()

	if err := exec.Command("open", path).Run(); err != nil {
		return err
	}
}
