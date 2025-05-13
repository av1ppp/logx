package handlerempty

import (
	"context"

	"github.com/av1ppp/logx"
)

type Handler struct {
}

func (self *Handler) Enabled(context.Context, logx.Level) bool {
	return false
}

func (self *Handler) Handle(context.Context, logx.Record) error {
	return nil
}

func (self *Handler) WithAttrs(attrs []logx.Attr) logx.Handler {
	return self
}

func (self *Handler) WithGroup(name string) logx.Handler {
	return self
}

func New() logx.Handler {
	return &Handler{}
}
