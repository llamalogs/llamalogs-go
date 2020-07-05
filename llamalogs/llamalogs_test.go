package llamalogs

import (
	"encoding/json"
	"sort"
	"testing"
)

func TestInit(t *testing.T) {
	Init(InitArgs{AccountKey: "key1", GraphName: "graphy1"})
	Log(LogArgs{Sender: "User", Receiver: "Web Server"})

	logList, _ := collectMessages()
	jsonValue, _ := json.Marshal(logList)
	expectedJSON := `[{"sender":"User","receiver":"Web Server","account":"key1","graph":"graphy1","count":1,"errors":0,"elapsed":0,"message":"","errorMessage":"","initialMessageCount":0}]`
	if string(jsonValue) != expectedJSON {
		t.Errorf("jsonValue is %s; expected is %s", jsonValue, expectedJSON)
	}
	// resetting initial init state
	Init(InitArgs{AccountKey: "", GraphName: ""})
}

func TestDisabledInit(t *testing.T) {
	Init(InitArgs{AccountKey: "key1", GraphName: "graphy1", Disabled: true})
	Log(LogArgs{Sender: "User", Receiver: "Web Server"})

	logList, _ := collectMessages()
	jsonValue, _ := json.Marshal(logList)
	expectedJSON := `[]`
	if string(jsonValue) != expectedJSON {
		t.Errorf("jsonValue is %s; expected is %s", jsonValue, expectedJSON)
	}
	// resetting initial init state
	Init(InitArgs{AccountKey: "", GraphName: "", Disabled: false})
}

func TestBasicLog(t *testing.T) {
	Log(LogArgs{Sender: "User", Receiver: "Web Server", Message: "request params: 123", GraphName: "g1", AccountKey: "ac1"})
	logList, _ := collectMessages()

	if len(logList) != 1 {
		t.Errorf("LogList length is %d; want 1", len(logList))
	}

	jsonValue, _ := json.Marshal(logList)
	expectedJSON := `[{"sender":"User","receiver":"Web Server","account":"ac1","graph":"g1","count":1,"errors":0,"elapsed":0,"message":"request params: 123","errorMessage":"","initialMessageCount":0}]`
	if string(jsonValue) != expectedJSON {
		t.Errorf("jsonValue is %s; expected is %s", jsonValue, expectedJSON)
	}
}

func TestMultipleLog(t *testing.T) {
	Log(LogArgs{Sender: "User", Receiver: "Web Server", Message: "request params: 123", GraphName: "g1", AccountKey: "ac1"})
	Log(LogArgs{Sender: "User", Receiver: "Web Server", Message: "other message", GraphName: "g1", AccountKey: "ac1"})
	Log(LogArgs{Sender: "Second Sender", Receiver: "Web Server", Message: "other message", GraphName: "g1", AccountKey: "ac1"})
	Log(LogArgs{Sender: "Second Sender", Receiver: "Web Server", IsError: true, Message: "error message", GraphName: "g1", AccountKey: "ac1"})
	logList, _ := collectMessages()
	// sorting just for test string, order doesnt matter for sending into proxy
	sort.Slice(logList, func(i, j int) bool {
		return logList[i].Sender < logList[j].Sender
	})

	if len(logList) != 2 {
		t.Errorf("LogList length is %d; want 1", len(logList))
	}

	jsonValue, _ := json.Marshal(logList)
	expectedJSON := `[{"sender":"Second Sender","receiver":"Web Server","account":"ac1","graph":"g1","count":2,"errors":1,"elapsed":0,"message":"other message","errorMessage":"error message","initialMessageCount":0},{"sender":"User","receiver":"Web Server","account":"ac1","graph":"g1","count":2,"errors":0,"elapsed":0,"message":"request params: 123","errorMessage":"","initialMessageCount":0}]`
	if string(jsonValue) != expectedJSON {
		t.Errorf("jsonValue is %s; expected is %s", jsonValue, expectedJSON)
	}
}
