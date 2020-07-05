package llamalogs

var globalGraphName = ""
var globalAccountKey = ""
var globalIsDevEnv = false

func Init(accountKey string, graphName string) {
	startTimer()
	globalGraphName = graphName
	globalAccountKey = accountKey
}

func SetDevEnv() {
	globalIsDevEnv = true
}

func Log(args LogArgs) {
	newLog := args.toLog()
	processLog(newLog)
}

func PointStat(args StatArgs) {
	newStat := args.toStat()
	newStat.kind = "point"
	processStat(newStat)
}

func AvgStat(args StatArgs) {
	newStat := args.toStat()
	newStat.kind = "average"
	processStat(newStat)
}

func MaxStat(args StatArgs) {
	newStat := args.toStat()
	newStat.kind = "max"
	processStat(newStat)
}

func ForceSend() {
	sendMessages()
}

func processStat(newStat stat) {
	if newStat.account == "" {
		newStat.account = globalAccountKey
	}

	if newStat.graph == "" {
		newStat.graph = globalGraphName
	}

	addStat(newStat)
}

func processLog(newLog logStruct) {
	if newLog.account == "" {
		newLog.account = globalAccountKey
	}

	if newLog.graph == "" {
		newLog.graph = globalGraphName
	}

	if newLog.sender == "" || newLog.receiver == "" || newLog.account == "" || newLog.graph == "" {
		return
	}

	addLog(newLog)
}
