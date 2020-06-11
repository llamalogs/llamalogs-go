package llamalogs

import "time"

type StatArgs struct {
	Component  string
	Name       string
	Value      float64
	AccountKey string
	GraphName  string
}

func (s StatArgs) toStat() stat {
	return stat{
		component: s.Component,
		name:      s.Name,
		value:     s.Value,
		account:   s.AccountKey,
		graph:     s.GraphName,
		timestamp: time.Now().UnixNano() / int64(time.Millisecond)}
}
