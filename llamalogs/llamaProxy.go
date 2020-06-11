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
	fmt.Printf("logList\n")
	fmt.Printf("%s\n", logList)

	for _, k := range existingStats {
		for _, v := range k {
			statList = append(statList, v.toJSONType())
		}
	}
	fmt.Printf("statList\n")
	fmt.Printf("%s\n", statList)

	if len(logList) == 0 && len(statList) == 0 {
		return
	}

	// 		if (len(log_list) or len(stat_list)):
	// 			try:
	// 				first_log = (len(log_list) and log_list[0]) or (len(stat_list) and stat_list[0])
	// 				account_key = first_log["account"]
	// 				requests.post(url + 'api/timedata', json = {"account_key": account_key, "time_logs": log_list, "time_stats": stat_list}, timeout=1)
	// 			except:
	// 				print('LlamaLogs Error; contacting llama logs server')

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

	fmt.Printf("json\n")
	fmt.Printf("%s\n", jsonValue)

	resp, err := http.Post("http://localhost:4000/api/timedata", "application/json", bytes.NewBuffer(jsonValue))
	if resp != nil {
		fmt.Printf("%s", resp)
	}
	if err != nil {
		fmt.Printf("%s", err)
	}

	// pd := make(map[string][]jsonLog)
	// pd["time_logs"] = logList
	// fmt.Printf("%s\n", logList)
	// jsonValue, err := json.Marshal(logList)
	// if err != nil {
	// 	log.Fatal("Cannot encode to JSON ", err)
	// }
	// fmt.Printf("json\n")
	// fmt.Printf("%s\n", jsonValue)

	// resp, err := http.Post("http://localhost:4000/api/time_logs", "application/json", bytes.NewBuffer(jsonValue))
	// if resp != nil {
	// 	fmt.Printf("%s", resp)
	// }
	// if err != nil {
	// 	fmt.Printf("%s", err)
	// }
}
