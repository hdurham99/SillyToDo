package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
)

var (
	idFlag     int
	addFlag    string
	deleteFlag string
	statusFlag string
)

func main() {
	fmt.Println("Silly To-Do App!")
	flag.IntVar(&idFlag, "id", -1, "id of task")
	flag.StringVar(&addFlag, "a", "", "adds task")
	flag.StringVar(&deleteFlag, "d", "", "deletes task")
	flag.StringVar(&statusFlag, "s", "", "outputs all tasks")
	flag.Parse()

	if addFlag != "" {
		add(addFlag, "output.csv")
	} else if deleteFlag != "" && idFlag != -1 {
		delete(idFlag)
	} else if statusFlag != "" {
		status()
	} else {
		fmt.Println("Invalid argument")
	}

}

// Will add the task
func add(task string, filename string) {
	taskID := strconv.Itoa(rand.IntN(15))
	writer, file, err := createCSVWriter(filename)
	if err != nil {
		fmt.Println("Error creating CSV writer: ", err)
		return
	}
	defer file.Close()
	header := []string{"ID", "Task"}
	writeCSVRecord(writer, header)
	record := []string{taskID, task}
	writeCSVRecord(writer, record)

	writer.Flush()
	if err := writer.Error(); err != nil {
		fmt.Println("Error flushing CSV writer: ", err)
	}
}

// Given a task ID, will delete that task
func delete(taskID int) {

}

// Will print all tasks that are still needing to be complete
func printTasks() {}

// Creates CSV File
func createCSVWriter(filename string) (*csv.Writer, *os.File, error) {
	f, err := os.Create(filename)
	if err != nil {
		return nil, nil, err
	}
	writer := csv.NewWriter(f)
	return writer, f, nil
}

func writeCSVRecord(writer *csv.Writer, record []string) {
	err := writer.Write(record)
	if err != nil {
		fmt.Println("Error writing record to CSV: ", err)
	}
}
