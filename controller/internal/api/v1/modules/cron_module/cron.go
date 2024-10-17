package cron_module

import (
	"fmt"

	cron "github.com/robfig/cron/v3"
)

var cronInstance *cron.Cron

func NewCronInstance() {
	fmt.Println("Starting new cron instance...")
	cronInstance = cron.New()
	cronInstance.Start()

	addAllFromDB()
}

func AddCron(data SCron) {
	cronInstance.AddFunc(data.At, func() { fmt.Println(data.Taskname) })
}

func ResetCrons() {
	entries := cronInstance.Entries()
	for _, entry := range entries {
		cronInstance.Remove(entry.ID)
	}
	addAllFromDB()
}

func addAllFromDB() {
	crons, err := GetAllCronsSQL()
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(crons.([]SCron)) > 0 {
		fmt.Printf("There is %v tasks to run !\n", len(crons.([]SCron)))
		for _, cron := range crons.([]SCron) {
			cronInstance.AddFunc(cron.At, func() { fmt.Println(cron.Taskname) })
		}
	}
}
