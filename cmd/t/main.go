package main

import (
	"os"

	"github.com/grplyler/tuego/pkg/tasks"
)

func main() {

	tm := tasks.NewTaskManager()
	tm.Load()

	// Parse Command line arguments
	if len(os.Args) == 1 {
		tm.WeeklySummary(tm.LatestWeek())
	}

}
