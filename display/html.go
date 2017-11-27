package display

import (
	"html/template"
	"os"
	"os/exec"

	"github.com/senorprogrammer/conf_check/modules"
)

type HtmlData struct {
	Keys        []string
	OutputPath  string
	RailsConfig *modules.RailsConfig
}

func NewHtmlData(railsConfig *modules.RailsConfig) *HtmlData {
	htmlData := HtmlData{
		Keys:        railsConfig.Keys(),
		OutputPath:  "./output/index.html",
		RailsConfig: railsConfig,
	}

	return &htmlData
}

func (htmlData *HtmlData) Render() {
	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		panic(err)
	}

	output, err := os.Create(htmlData.OutputPath)
	defer output.Close()
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(output, htmlData)
	if err != nil {
		panic(err)
	}
}

func (htmlData *HtmlData) Show() {
	if err := exec.Command("open", htmlData.OutputPath).Run(); err != nil {
		panic(err)
	}
}
