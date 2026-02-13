package handlercolor

import (
	"strings"
)

// Prefix prepends a colored prefix to msg.
func Prefix(prefix string, msg ...string) string {
	if len(msg) == 0 {
		return stylePrefix.Render(prefix)
	}
	return stylePrefix.Render(prefix) + " " + strings.Join(msg, " ")
}
