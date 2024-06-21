package internal

import (
	"github.com/dronestock/chat/internal/internal/config"
	"github.com/dronestock/chat/internal/internal/step"
	"github.com/dronestock/drone"
	"github.com/larksuite/oapi-sdk-go/v3"
)

type plugin struct {
	drone.Base

	Upload  *config.Upload   `default:"${UPLOAD={}}" json:"upload,omitempty"`
	Uploads []*config.Upload `default:"${UPLOADS}" json:"uploads,omitempty"`

	// 飞书
	Lark *config.Lark `default:"${LARK=${FEISHU}}" json:"lark,omitempty"`

	lark *lark.Client
}

func New() drone.Plugin {
	return new(plugin)
}

func (p *plugin) Config() drone.Config {
	return p
}

func (p *plugin) Steps() drone.Steps {
	return drone.Steps{
		drone.NewStep(step.NewTodo()).Name("样例").Build(),
	}
}

func (p *plugin) Setup() (err error) {
	if p.Upload.Checked() {
		p.Uploads = append(p.Uploads, p.Upload)
	}

	if nil != p.Lark {
		p.lark = lark.NewClient(p.Lark.Id, p.Lark.Secret)
	}

	return
}
