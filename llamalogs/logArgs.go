package llamalogs

import "time"

type LogArgs struct {
	Sender     string
	Receiver   string
	Message    string
	IsError    bool
	AccountKey string
	GraphName  string
}

func (l LogArgs) toLog() logStruct {
	return logStruct{
		sender:           l.Sender,
		receiver:         l.Receiver,
		timestamp:        time.Now().UnixNano() / int64(time.Millisecond),
		log:              l.Message,
		isInitialMessage: true,
		account:          l.AccountKey,
		graph:            l.GraphName,
		isError:          l.IsError,
		elapsed:          0}
}
