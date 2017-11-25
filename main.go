package main

import (
	"flag"
	"fmt"

	"github.com/senorprogrammer/conf_check/modules"
)

func main() {
	path := flag.String("dir", "./", "Path to Rails application")
	flag.Parse()

	fmt.Println("Running conf_check...")

	railsConf := modules.NewRailsConfig(path)
	railsConf.Load(path)
	railsConf.Check()
}
