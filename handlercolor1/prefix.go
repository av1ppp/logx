package handlercolor1

import (
	"strings"
)

// Prefix prepends a colored prefix to msg.
func Prefix(prefix string, msg ...string) string {
	if len(msg) == 0 {
		return colorPrefix.Sprint(prefix)
	}
	return colorPrefix.Sprint(prefix) + " " + strings.Join(msg, " ")
}
