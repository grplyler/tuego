package main

import (
	"os"
)

func main() {

	tm := NewTaskManager()
	tm.Load()

	// Parse Command line argumnets
	if len(os.Args) == 1 {
		tm.WeeklySummary(tm.LatestWeek())
	}

}
