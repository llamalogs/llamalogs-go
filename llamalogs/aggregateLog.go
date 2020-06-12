package llamalogs

type aggregatedLog struct {
	sender              string
	receiver            string
	account             string
	count               int
	errors              int
	elapsed             int
	log                 string
	errorLog            string
	initialMessageCount int
	graph               string
}

type jsonLog struct {
	Sender              string `json:"sender"`
	Receiver            string `json:"receiver"`
	Account             string `json:"account"`
	Graph               string `json:"graph"`
	Count               int    `json:"count"`
	Errors              int    `json:"errors"`
	Elapsed             int    `json:"elapsed"`
	Log                 string `json:"log"`
	ErrorLog            string `json:"errorLog"`
	InitialMessageCount int    `json:"initialMessageCount"`
}

func (ag aggregatedLog) toJSONType() jsonLog {
	return jsonLog{
		Sender:              ag.sender,
		Receiver:            ag.receiver,
		Account:             ag.account,
		Count:               ag.count,
		Errors:              ag.errors,
		Elapsed:             ag.elapsed,
		Log:                 ag.log,
		ErrorLog:            ag.errorLog,
		InitialMessageCount: ag.initialMessageCount,
		Graph:               ag.graph}
}
