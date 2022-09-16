package system

import (
	"fmt"
	"github.com/araddon/dateparse"
	"log"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"time"
)

func System(cmd string, arg ...string) {
	out, err := exec.Command(cmd, arg...).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}
func CheckFiles() {
	_, err := File2lines("./config")
	if err != nil {
		AddConfigDefault()
		os.Exit(0)
	}
	_, err2 := os.Stat(OwlFolderLogs)
	if os.IsNotExist(err2) {
		Run("mkdir " + OwlFolderLogs)
	}
	_, err4 := File2lines(OwlAccesslog)
	if err4 != nil {
		createFile(OwlAccesslog)
		_, _ = RunString("chmod -R 777 " + OwlAccesslog)
	}
}

func Config(data string) string {
	lines, err := File2lines("./config")
	value := ""
	if err != nil {
		fmt.Println("The configuration file could not be found")
	} else {
		// --- Extract the variables from the configuration file
		for _, line := range lines {
			if strings.Contains(line, data) {
				cut := strings.Split(line, "=")
				value += cut[1]
			}
		}
	}
	// -------------------------------------------
	return value
}

func SubtractDates(date string) int {
	a, _ := strconv.ParseInt(date, 10, 64)
	ta := time.Unix(a, 0)
	loc, _ := time.LoadLocation("UTC")
	time.Local = loc
	parse, _ := dateparse.ParseLocal(ta.String())
	return int(time.Now().Sub(parse).Seconds())
}

func Contains(s []string, searcher string) bool {
	i := sort.SearchStrings(s, searcher)
	return i < len(s) && s[i] == searcher
}

// AddConfigDefault --Add the default configuration to the configuration file
func AddConfigDefault() {
	Run("touch ./config")
	_ = AppendStrFile("./config", "\n")
	_ = AppendStrFile("./config", "#--	 System configuration file --\n")
	_ = AppendStrFile("./config", "#--  Created by Edgar Javier akosej9208@gmail.com  --\n")
	_ = AppendStrFile("./config", "#--  System configuration file   --\n")
	_ = AppendStrFile("./config", "#-- Path necessary files\n")
	_ = AppendStrFile("./config", "path.AccessLog=./access.log\n")
	_ = AppendStrFile("./config", "#--path where the logs will be saved\n")
	_ = AppendStrFile("./config", "folder.salve_logs=./salva_logs\n")
	_ = AppendStrFile("./config", "#-- Interface server\n")
	_ = AppendStrFile("./config", "interface.server=eth0\n")
	_ = AppendStrFile("./config", "#-- Default quota for users 50 mb 1mb=1048576 Bytes\n")
	_ = AppendStrFile("./config", "default.quota=52428800\n")
	_ = AppendStrFile("./config", "#-- Exposed server interface to clients\n")
	_ = AppendStrFile("./config", "interface.server=eth0\n")
	_ = AppendStrFile("./config", "#-- Exposed squid port\n")
	_ = AppendStrFile("./config", "squid.port=3128\n")
	_ = AppendStrFile("./config", "#-- Exposed squid port ssl\n")
	_ = AppendStrFile("./config", "squid.port=3129\n")
	_ = AppendStrFile("./config", "#-------Hours in which the system will restart the quota and rotate the logs  ---------------\n")
	_ = AppendStrFile("./config", "jobs.restart=08:00:00 10:00:00 11:00:00 11:30:00 12:00:00 13:00:00 14:00:00 15:00:00 16:00:00\n")
	_ = AppendStrFile("./config", "#--Redis server Ipaddress\n")
	_ = AppendStrFile("./config", "redis.ip=127.0.0.1\n")
	_ = AppendStrFile("./config", "#--Redis server Password\n")
	_ = AppendStrFile("./config", "redis.pass=pass\n")
	_ = AppendStrFile("./config", "#--Redis server Port\n")
	_ = AppendStrFile("./config", "redis.port=6379\n")
	fmt.Println("Restart the OWL system, the necessary files have been created.")
}

// Ct --Current time
func Ct() time.Time {
	return time.Now()
}
//
//func ResetFile(path string) {
//	_ = os.Remove(path)
//	createFile(path)
//	_, _ = RunString("chmod -R 777 " + path)
//}
