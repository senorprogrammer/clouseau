package modules

import (
	"io/ioutil"
	"os"

	"github.com/stretchr/powerwalk"
)

type Checkbox struct {
	FileChecks []Checkable
	Path       string
}

/* -------------------- Public Functions -------------------- */

func (checkbox *Checkbox) Append(checkable Checkable) {
	checkbox.FileChecks = append(checkbox.FileChecks, checkable)
}

func (checkbox *Checkbox) Run() {
	powerwalk.Walk(checkbox.Path, func(path string, info os.FileInfo, err error) error {
		file, _ := os.Open(path)
		defer file.Close()

		data, err := ioutil.ReadAll(file)
		if err == nil {
			for _, checkable := range checkbox.FileChecks {
				checkable.Parse(string(data), file.Name())
			}
		}
		return nil
	})

	/* Clean up the search results (TODO: move this out of ConfigChecker) */
	for _, checkable := range checkbox.FileChecks {
		checkable.Sanitize()
	}
}
