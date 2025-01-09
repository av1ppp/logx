package logx

import (
	"github.com/joomcode/errorx"
)

var (
	root         = errorx.NewNamespace("logx")
	commonErrors = root.NewType("common")
)
