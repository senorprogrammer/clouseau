package display

import (
	"html/template"
	"os"
	"os/exec"

	"github.com/senorprogrammer/conf_check/modules"
)

type HtmlData struct {
	Keys               []string
	OutputPath         string
	RailsConfigScanner *modules.RailsConfigScanner
}

func NewHtmlData(railsConfig *modules.RailsConfigScanner) *HtmlData {
	data := HtmlData{
		Keys:               railsConfig.Keys(),
		OutputPath:         "./output/index.html",
		RailsConfigScanner: railsConfig,
	}

	return &data
}

func (data *HtmlData) Render() {
	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		panic(err)
	}

	output, err := os.Create(data.OutputPath)
	defer output.Close()
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(output, data)
	if err != nil {
		panic(err)
	}
}

func (data *HtmlData) Show() {
	if err := exec.Command("open", data.OutputPath).Run(); err != nil {
		panic(err)
	}
}
