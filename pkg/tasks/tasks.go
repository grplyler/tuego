package tasks

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/grplyler/tuego/pkg/format"
)

type TaskManager struct {
	Data [][]string
}

func NewTaskManager() *TaskManager {
	tm := new(TaskManager)
	tm.Data = make([][]string, 10)
	return tm
}

func (tm *TaskManager) Load() {
	csvfile, err := os.Open("tue.csv")
	if err != nil {
		log.Print("Error opening data file", err)
	}

	reader := csv.NewReader(csvfile)

	if _, err := reader.Read(); err != nil {
		panic(err)
	}
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error reading CSV data", err)
	}

	for _, rec := range records {
		fmt.Println(rec)
	}

	tm.Data = records
	fmt.Println("Data loaded")

}

func (tm *TaskManager) Save() {
	file, err := os.Create("result.csv")
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range tm.Data {
		err := writer.Write(value)
		checkError("Cannot write to file", err)
	}
}

func (tm *TaskManager) Progress(week string) string {
	total := 0
	done := 0
	for _, row := range tm.Data {
		iprog, err := strconv.Atoi(row[3])
		checkError("Could not parse progress int", err)
		done += iprog
		total += 100
	}

	fmt.Println(done)
	fmt.Println(total)
	progress := float64(done) / float64(total) * 100
	result := fmt.Sprintf("Progress %.2f", progress)

	return result
}

func (tm *TaskManager) LatestWeek() string {

	var latest_week int64 = 0
	for _, todo := range tm.Data {
		current_week, err := strconv.ParseInt(todo[1], 10, 32)
		checkError("Error parsing week number from data", err)

		if current_week > latest_week {
			latest_week = current_week
		}

	}

	return fmt.Sprintf("%v", latest_week)
}

func (tm *TaskManager) WeeklySummary(week string) {

	// total := 0
	// done := 0
	current_class := ""
	last_class := ""
	for index, todo := range tm.Data {
		current_class = todo[0]

		// Print Class Header
		if current_class != last_class {
			name := fmt.Sprintf("[ %v ]", todo[0])
			line := fmt.Sprintf("+----+----------+%v", format.Filll(name, 45, "-", "", ""))
			fmt.Println(line)
			last_class = current_class
		}

		prog, _ := strconv.ParseFloat(todo[3], 10)
		prog_fmt := format.ProgressBar(prog, 10, "█", "", "|")
		title := todo[2]
		index_fmt := format.Filll(strconv.Itoa(index), 4, " ", "|", "|")
		line := fmt.Sprintf("%s%s%s", index_fmt, prog_fmt, title)
		fmt.Println(line)

	}
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(fmt.Sprintf("%s:\n\t", message), err)
	}
}
