package main

import (
	"copyLog/database"
	"copyLog/system"
	"github.com/hpcloud/tail"
	"sort"
	"strings"
)

func main() {
	// -- Check that the necessary files exist
	system.CheckFiles()
	// --Connect to the mysql server
	database.Connect()

	t, _ := tail.TailFile(system.OwlAccesslog, tail.Config{Follow: true})

	sort.Strings(system.ConnectType)

	for line := range t.Lines {
		segment := strings.Fields(line.Text)
		//-- Write DB
		database.CreateLogs(segment)
	}
}
