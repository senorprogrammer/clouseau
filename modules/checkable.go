package modules

import ()

type Checkable interface {
	Keys() []string
	Len() int
	Name() string
	Parse(data string, fileName string)
	Sanitize()
}
