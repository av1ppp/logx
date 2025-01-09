package rotation

import (
	"github.com/joomcode/errorx"
)

var (
	root         = errorx.NewNamespace("logx_rotation")
	commonErrors = root.NewType("common")
)
