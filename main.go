package main

import (
	"copyLog/database"
	"copyLog/system"
	"github.com/hpcloud/tail"
	"sort"
	"strings"
)

func main() {
	//-------Check that the necessary files exist
	system.CheckFiles()
	//-------Connect to the redis server
	database.Connect()
	//-------JOBS RESTARTS SYSTEM
	//-------Read real-time access.log from squid
	t, _ := tail.TailFile(system.OwlAccesslog, tail.Config{Follow: true})

	sort.Strings(system.ConnectType)

	for line := range t.Lines {
		//time duration client_address result_code bytes request_method url rfc931 hierarchy_code type
		// 0        1       2               3       4       5           6       7           8       9
		segment := strings.Fields(line.Text)
		//-- Write DB
		database.CreateLogs(segment)
	}
}
