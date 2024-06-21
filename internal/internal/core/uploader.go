package core

import (
	"os"
)

type Uploader interface {
	Upload(file *os.File) error
}
