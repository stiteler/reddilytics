package main

import (
	"github.com/steaz/reddilytics/model"
	"github.com/steaz/reddilytics/update"
	"time"
)

func main() {
	for {
		err := model.SetupDB()
		if err != nil {
			panic(err)
		}

		update.UpdateAll()

		//once an hour?
		time.Sleep(time.Hour*24)
	}
}
