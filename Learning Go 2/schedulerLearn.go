package main

import (
	"fmt"
	"time"

	"github.com/carlescere/scheduler"
	"net/http"
)

func main() {
	job := func() {
		t := time.Now()
		fmt.Println("Time's up! @", t.UTC())
	}
	// Run every 2 seconds but not now.
	scheduler.Every(2).Seconds().NotImmediately().Run(job)

	// Run now and every X.
	//scheduler.Every(5).Minutes().Run(job)
	//scheduler.Every().Day().Run(job)
	//scheduler.Every().Monday().At("08:30").Run(job)

	fmt.Println("other task")
	time.Sleep(time.Second*3)
	fmt.Println("other task 2")

	// Keep the program from not exiting.
	//runtime.Goexit()
	http.ListenAndServe(":8000" , nil)
}