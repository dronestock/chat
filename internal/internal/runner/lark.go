package runner

import (
	"os"

	"github.com/dronestock/chat/internal/internal/config"
	"github.com/dronestock/chat/internal/internal/core"
	"github.com/goexl/gox"
	"github.com/larksuite/oapi-sdk-go/v3"
)

var _ core.Runner = (*Lark)(nil)

type Lark struct {
	client *lark.Client
	_      gox.CannotCopy
}

func NewLark(config *config.Lark) *Lark {
	return &Lark{
		client: lark.NewClient(config.Id, config.Secret),
	}
}

func (l *Lark) Upload(name string, file *os.File) (err error) {
	return
}

func (l *Lark) Notify() (err error) {
	return
}
