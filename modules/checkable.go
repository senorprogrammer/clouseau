package modules

import ()

type Checkable interface {
	Keys() []string
	Len() int
	Name() string
	Parse(line string, fileName string)
	Sanitize()
}
