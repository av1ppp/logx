package handlerdev

import "github.com/av1ppp/logx"

type attributes []logx.Attr

func (a attributes) Len() int      { return len(a) }
func (a attributes) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a attributes) Less(i, j int) bool {
	if a[i].Value.Kind() == logx.KindGroup && a[j].Value.Kind() != logx.KindGroup {
		return false
	} else if a[i].Value.Kind() != logx.KindGroup && a[j].Value.Kind() == logx.KindGroup {
		return true
	}

	return a[i].Key < a[j].Key
}
