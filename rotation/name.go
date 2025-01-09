package rotation

import (
	"path/filepath"
	"strings"
	"time"
)

func makeLogName(prefix string) string {
	return prefix + ".log"
}

func makeZipName(prefix string) string {
	return prefix + "-" + time.Now().UTC().Format(zipTimeLayout) + zipExt
}

func extractCreateTimeFromZipName(name string) (time.Time, bool) {
	s := strings.TrimSuffix(filepath.Base(name), zipExt)
	if len(s) < 19 {
		return time.Time{}, false
	}

	t, err := time.Parse(zipTimeLayout, s[len(s)-19:])
	if err != nil {
		return time.Time{}, false
	}

	return t, true
}
