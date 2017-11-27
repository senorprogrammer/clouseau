package display

import (
	"os"
	"sort"

	// "github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/senorprogrammer/conf_check/modules"
)

type Table struct{}

func (table *Table) Render(railsConfig *modules.RailsConfig) {
	keys := railsConfig.Keys()
	sort.Strings(keys)

	tableWtr := tablewriter.NewWriter(os.Stdout)

	headerArr := []string{"key"}
	for _, configFile := range railsConfig.ConfigFiles {
		headerArr = append(headerArr, configFile.Name)
	}

	for _, key := range keys {
		arr := []string{key}

		for _, configFile := range railsConfig.ConfigFiles {
			arr = append(arr, configFile.Entries[key].Value)
		}

		tableWtr.Append(arr)
	}

	tableWtr.SetHeader(headerArr)
	tableWtr.Render()
}
