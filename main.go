package main

import (
	"flag"
	"fmt"
)

var (
	idFlag     int
	addFlag    string
	deleteFlag string
	statusFlag string
)

func main() {
	fmt.Println("Silly To-Do App!")
	flag.IntVar(&idFlag, "id", 0, "id of task")
	flag.StringVar(&addFlag, "a", "", "adds task")
	flag.StringVar(&deleteFlag, "d", "", "deletes task")
	flag.StringVar(&statusFlag, "s", "", "outputs all tasks")
	flag.Parse()

	fmt.Println("idFlag value is: ", idFlag)
	fmt.Println("addFlag value is: ", addFlag)
	fmt.Println("deletFlag value is: ", deleteFlag)
	fmt.Println("statusFlag value is: ", statusFlag)
}
