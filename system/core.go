package system

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
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
	_, err4 := File2lines(Accesslog)
	if err4 != nil {
		createFile(Accesslog)
		_, _ = RunString("chmod -R 777 " + Accesslog)
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

// AddConfigDefault --Add the default configuration to the configuration file
func AddConfigDefault() {
	Run("touch ./config")
	_ = AppendStrFile("./config", "\n")
	_ = AppendStrFile("./config", "#--  System configuration file --\n")
	_ = AppendStrFile("./config", "#--  Created by Edgar Javier akosej9208@gmail.com  --\n")
	_ = AppendStrFile("./config", "#----------------------------------------------------\n")
	_ = AppendStrFile("./config", "#-- Path necessary files\n")
	_ = AppendStrFile("./config", "path.AccessLog=./access.log\n")
	_ = AppendStrFile("./config", "#----------------------------------------------------\n")
	_ = AppendStrFile("./config", "#--Mysql server Ipaddress\n")
	_ = AppendStrFile("./config", "mysql.ip=127.0.0.1\n")
	_ = AppendStrFile("./config", "#--Mysql server Username\n")
	_ = AppendStrFile("./config", "myslq.user=root\n")
	_ = AppendStrFile("./config", "#--Mysql server Password\n")
	_ = AppendStrFile("./config", "myslq.pass=pass\n")
	_ = AppendStrFile("./config", "#--Myslq server Port\n")
	_ = AppendStrFile("./config", "myslq.port=3306\n")
	_ = AppendStrFile("./config", "#--Myslq Databese\n")
	_ = AppendStrFile("./config", "myslq.db=reports\n")
	fmt.Println("Restart the binary, the necessary files have been created.")
}
