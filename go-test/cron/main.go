package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New(cron.WithSeconds(), cron.WithLocation(time.UTC)) // Use seconds if needed

	// Schedule: At 00:00 on the 1st of every month
	_, err := c.AddFunc("0 0 0 1 * *", func() {
		fmt.Println("Running monthly task at", time.Now())
		// Your job logic here
	})

	if err != nil {
		panic(err)
	}

	c.Start()

	// Keep the program running
	select {}
}
