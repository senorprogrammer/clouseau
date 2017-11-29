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
	EnvVarChecker      *modules.EnvVarChecker
	RailsConfigChecker *modules.RailsConfigChecker
}

func NewHtmlData(envVarChecker *modules.EnvVarChecker, railsConfChecker *modules.RailsConfigChecker) *HtmlData {
	data := HtmlData{
		EnvVarChecker:      envVarChecker,
		RailsConfigChecker: railsConfChecker,
		OutputPath:         "./output/index.html",
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
