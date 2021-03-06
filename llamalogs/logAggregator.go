package llamalogs

import (
	"fmt"
	"sync"
	"time"
)

type logMap = map[string]map[string]*aggregatedLog
type statMap = map[string]map[string]*stat

var logMutex = &sync.Mutex{}

var aggregateLogs = make(logMap)
var aggregateStats = make(statMap)
var timerStarted = false

func startTimer() {
	if timerStarted {
		return
	}
	// sending in data after 5 seconds for the first time for early results
	go func() {
		time.Sleep(5 * time.Second)
		go sendMessages()
	}()

	ticker := time.NewTicker(59500 * time.Millisecond)

	go func() {
		for {
			select {
			case _ = <-ticker.C:
				go sendMessages()
			}
		}
	}()

	timerStarted = true
}

func getAndClearLogs() (logMap, statMap) {
	logMutex.Lock()
	existingLogs := aggregateLogs
	existingStats := aggregateStats
	aggregateLogs = make(logMap)
	aggregateStats = make(statMap)
	logMutex.Unlock()
	return existingLogs, existingStats
}

func addLog(newLog logStruct) {
	logMutex.Lock()

	if _, found := aggregateLogs[newLog.sender]; !found {
		aggregateLogs[newLog.sender] = make(map[string]*aggregatedLog)
	}

	if _, found := aggregateLogs[newLog.sender][newLog.receiver]; !found {
		newAgg := newLog.toAggregate()
		aggregateLogs[newLog.sender][newLog.receiver] = &newAgg
	}

	existing := aggregateLogs[newLog.sender][newLog.receiver]

	if newLog.isError {
		existing.errors = existing.errors + 1
	}

	// if (log.elapsed):
	// 	prev_amount = working_ob.elapsed * working_ob.elapsedCount
	// 	working_ob.elapsed = (prev_amount + log.elapsed) / (working_ob.total + 1)
	// 	working_ob.elapsedCount = working_ob.elapsedCount + 1
	// if (log.initialMessage):
	// 	working_ob.initialMessageCount = working_ob.initialMessageCount + 1

	existing.count = existing.count + 1
	if existing.message == "" && !newLog.isError {
		existing.message = newLog.message
	}
	if existing.errorMessage == "" && newLog.isError {
		existing.errorMessage = newLog.message
	}

	// aggregateLogs[newLog.sender][newLog.receiver] = existing

	if globalIsDevEnv {
		fmt.Println(aggregateLogs)
	}

	logMutex.Unlock()
}

func addStat(newStat stat) {
	logMutex.Lock()

	if newStat.kind == "point" {
		if componentMap, found := aggregateStats[newStat.component]; found {
			componentMap[newStat.name] = &newStat
		} else {
			aggregateStats[newStat.component] = make(map[string]*stat)
			aggregateStats[newStat.component][newStat.name] = &newStat
		}
	}

	if newStat.kind == "average" {
		addStatAvg(newStat)
	}
	if newStat.kind == "max" {
		addStatMax(newStat)
	}

	logMutex.Unlock()
}

func addStatAvg(newStat stat) {
	if _, found := aggregateStats[newStat.component]; !found {
		aggregateStats[newStat.component] = make(map[string]*stat)
	}

	if _, found := aggregateStats[newStat.component][newStat.name]; !found {
		aggregateStats[newStat.component][newStat.name] = &newStat
		newStat.count = 0
	}

	existing := aggregateStats[newStat.component][newStat.name]
	existing.value = existing.value + newStat.value
	existing.count = existing.count + 1
}

func addStatMax(newStat stat) {
	if _, found := aggregateStats[newStat.component]; !found {
		aggregateStats[newStat.component] = make(map[string]*stat)
	}

	if _, found := aggregateStats[newStat.component][newStat.name]; !found {
		aggregateStats[newStat.component][newStat.name] = &newStat
	}
	existing := aggregateStats[newStat.component][newStat.name]
	if newStat.value > existing.value {
		existing.value = newStat.value
	}
}
