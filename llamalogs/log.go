package llamalogs

type logStruct struct {
	sender           string
	receiver         string
	timestamp        int64
	log              string
	isInitialMessage bool
	account          string
	graph            string
	isError          bool
	elapsed          int
}

func (l logStruct) toAggregate() aggregatedLog {
	return aggregatedLog{
		sender:   l.sender,
		receiver: l.receiver,
		account:  l.account,
		graph:    l.graph}
}
