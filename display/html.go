package display

import (
	"html/template"
	"os"
	"os/exec"
	// "strings"
	"path/filepath"

	"github.com/gobuffalo/packr"
	"github.com/senorprogrammer/clouseau/modules"
)

type HtmlData struct {
	Keys               []string
	OutputDir          string
	OutputFile         string
	EnvVarChecker      *modules.ConfigChecker
	ConfigChecker      *modules.ConfigChecker
	FigaroChecker      *modules.ConfigChecker
	RailsConfigChecker *modules.RailsConfigChecker
}

func NewHtmlData(envVarChecker, configChecker, figaroChecker *modules.ConfigChecker, railsConfChecker *modules.RailsConfigChecker) *HtmlData {
	data := HtmlData{
		EnvVarChecker:      envVarChecker,
		ConfigChecker:      configChecker,
		FigaroChecker:      figaroChecker,
		RailsConfigChecker: railsConfChecker,
		OutputDir:          "clouseau",
		OutputFile:         "index.html",
	}

	return &data
}

/* -------------------- Public Functions -------------------- */

func (data *HtmlData) Render() {
	box := packr.NewBox("../templates")
	tmpl, err := template.New("index").Parse(box.String("index.html"))
	if err != nil {
		panic(err)
	}

	output, err := os.Create(data.outputPath())
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
	if err := exec.Command("open", data.outputPath()).Run(); err != nil {
		panic(err)
	}
}

/* -------------------- Private Functions -------------------- */

func (data *HtmlData) outputPath() string {
	path := filepath.Join(data.OutputDir, data.OutputFile)
	err := os.MkdirAll(data.OutputDir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	return path
}
