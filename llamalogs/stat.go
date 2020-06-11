package llamalogs

type stat struct {
	component string
	name      string
	kind      string
	value     float64
	account   string
	graph     string
	timestamp int64
	count     int
}

type jsonStat struct {
	Component string  `json:"component"`
	Name      string  `json:"name"`
	Kind      string  `json:"type"`
	Value     float64 `json:"value"`
	Account   string  `json:"account"`
	Graph     string  `json:"graph"`
	Timestamp int64   `json:"timestamp"`
	Count     int     `json:"count"`
}

func (s stat) toJSONType() jsonStat {
	return jsonStat{
		Component: s.component,
		Name:      s.name,
		Kind:      s.kind,
		Value:     s.value,
		Account:   s.account,
		Graph:     s.graph,
		Timestamp: s.timestamp,
		Count:     s.count}
}
