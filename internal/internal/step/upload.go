package step

import (
	"context"

	"github.com/dronestock/chat/internal/internal/config"
	lark "github.com/larksuite/oapi-sdk-go/v3"
)

type Upload struct {
	uploads []*config.Upload
	lark    *lark.Client
}

func NewUpload(uploads []*config.Upload, lark *lark.Client) *Upload {
	return &Upload{
		uploads: uploads,
		lark:    lark,
	}
}

func (u *Upload) Runnable() bool {
	return 0 != len(u.uploads)
}

func (u *Upload) Run(_ *context.Context) (err error) {
	return
}

func (u *Upload) toLark(_ *context.Context) (err error) {
	return
}
