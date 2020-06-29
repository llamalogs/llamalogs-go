package llamalogs

type aggregatedLog struct {
	sender              string
	receiver            string
	account             string
	count               int
	errors              int
	elapsed             int
	message             string
	errorMessage        string
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
	Message             string `json:"message"`
	ErrorMessage        string `json:"errorMessage"`
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
		Message:             ag.message,
		ErrorMessage:        ag.errorMessage,
		InitialMessageCount: ag.initialMessageCount,
		Graph:               ag.graph}
}
