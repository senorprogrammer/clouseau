package display

import (
	"os"
	"sort"

	"github.com/olekukonko/tablewriter"
	"github.com/senorprogrammer/conf_check/modules"
)

type TableData struct {
	Keys               []string
	RailsConfigChecker *modules.RailsConfigChecker
}

func NewTableData(checker *modules.RailsConfigChecker) *TableData {
	data := TableData{
		RailsConfigChecker: checker,
	}

	return &data
}

func (data *TableData) Render() {
	keys := data.RailsConfigChecker.Keys()
	sort.Strings(keys)

	tableWtr := tablewriter.NewWriter(os.Stdout)

	headerArr := []string{"key"}
	for _, configFile := range data.RailsConfigChecker.ConfigFiles {
		headerArr = append(headerArr, configFile.Name)
	}

	for _, key := range keys {
		arr := []string{key}

		for _, configFile := range data.RailsConfigChecker.ConfigFiles {
			arr = append(arr, configFile.Entries[key].Value)
		}

		tableWtr.Append(arr)
	}

	tableWtr.SetHeader(headerArr)
	tableWtr.Render()
}
