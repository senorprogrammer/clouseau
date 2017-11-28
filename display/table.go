package display

import (
	"os"
	"sort"

	// "github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/senorprogrammer/conf_check/modules"
)

type TableData struct {
	Keys               []string
	RailsConfigScanner *modules.RailsConfigScanner
}

func NewTableData(railsConfig *modules.RailsConfigScanner) *TableData {
	data := TableData{
		RailsConfigScanner: railsConfig,
	}

	return &data
}

func (data *TableData) Render() {
	keys := data.RailsConfigScanner.Keys()
	sort.Strings(keys)

	tableWtr := tablewriter.NewWriter(os.Stdout)

	headerArr := []string{"key"}
	for _, configFile := range data.RailsConfigScanner.ConfigFiles {
		headerArr = append(headerArr, configFile.Name)
	}

	for _, key := range keys {
		arr := []string{key}

		for _, configFile := range data.RailsConfigScanner.ConfigFiles {
			arr = append(arr, configFile.Entries[key].Value)
		}

		tableWtr.Append(arr)
	}

	tableWtr.SetHeader(headerArr)
	tableWtr.Render()
}
