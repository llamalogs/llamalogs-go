package llamalogs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func sendMessages() {
	existingLogs, existingStats := getAndClearLogs()

	logList := []jsonLog{}
	statList := []jsonStat{}

	for _, k := range existingLogs {
		for _, v := range k {
			logList = append(logList, v.toJSONType())
		}
	}

	for _, k := range existingStats {
		for _, v := range k {
			statList = append(statList, v.toJSONType())
		}
	}

	if len(logList) == 0 && len(statList) == 0 {
		return
	}

	type TimeReq struct {
		AccountKey string     `json:"account_key"`
		TimeLogs   []jsonLog  `json:"time_logs"`
		TimeStats  []jsonStat `json:"time_stats"`
	}

	newReq := TimeReq{}

	if len(logList) > 0 {
		newReq.AccountKey = logList[0].Account
	}
	if len(statList) > 0 {
		newReq.AccountKey = statList[0].Account
	}

	newReq.TimeLogs = logList
	newReq.TimeStats = statList

	jsonValue, err := json.Marshal(newReq)
	if err != nil {
		log.Fatal("Cannot encode to JSON ", err)
	}

	resp, err := http.Post("http://llamalogs.com/api/v0/timedata", "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		fmt.Printf("LlamaLogs Error; Error sending data - %s", err)
	}
}
