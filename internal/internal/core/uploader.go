package core

import (
	"os"
)

type Uploader interface {
	Upload(name string, file *os.File) error
}
