package llamalogs

var globalGraphName = ""
var globalAccountKey = ""
var globalIsDevEnv = false
var globalIsDisabled = false

func Init(args InitArgs) {
	globalGraphName = args.GraphName
	globalAccountKey = args.AccountKey
	globalIsDevEnv = args.IsDevEnv
	globalIsDisabled = args.Disabled

	if globalIsDisabled != true {
		startTimer()
	}
}

func Log(args LogArgs) {
	if globalIsDisabled {
		return
	}
	newLog := args.toLog()
	processLog(newLog)
}

func PointStat(args StatArgs) {
	if globalIsDisabled {
		return
	}
	newStat := args.toStat()
	newStat.kind = "point"
	processStat(newStat)
}

func AvgStat(args StatArgs) {
	if globalIsDisabled {
		return
	}
	newStat := args.toStat()
	newStat.kind = "average"
	processStat(newStat)
}

func MaxStat(args StatArgs) {
	if globalIsDisabled {
		return
	}
	newStat := args.toStat()
	newStat.kind = "max"
	processStat(newStat)
}

func ForceSend() {
	if globalIsDisabled {
		return
	}
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
