package database

import (
	"copyLog/models"
	"copyLog/system"
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open(system.MysqlCredentials), &gorm.Config{})
	if err != nil {
		fmt.Println("Could not connect to the database")
	}
	DB = connection
	_ = connection.AutoMigrate(&models.Report{})
}

func conn() *gorm.DB {
	connection, err := gorm.Open(mysql.Open(system.MysqlCredentials), &gorm.Config{})
	if err != nil {
		fmt.Println("Could not connect to the database")
	}
	return connection
}

func CreateLogs(date []string) bool {
	var r models.Report
	//1663246936.913 5815 192.168.0.28 TCP_TUNNEL/200 1202 CONNECT meet-jit-si-uk-london-1-local-3285-jvb-42-89-117.jitsi.net:443 liubel FIRSTUP_PARENT/10.60.7.117 -
	//time duration client_address result_code bytes request_method url rfc931 hierarchy_code type
	// 0        1       2               3       4       5           6       7           8       9
	//Elapsed    sql.NullInt64  `db:"elapsed"`
	//Remotehost sql.NullString `db:"remotehost"`
	//CodeStatus sql.NullString `db:"code_status"`
	//Bytes      sql.NullInt64  `db:"bytes"`
	//Method     sql.NullString `db:"method"`
	//URL        sql.NullString `db:"url"`
	//Rfc931     sql.NullString `db:"rfc931"`
	//PeerStatus sql.NullString `db:"peer_status"`
	//Type       sql.NullString `db:"type"`

	i, _ := strconv.ParseInt(strings.Split(date[0], ".")[0], 10, 64)
	tm := time.Unix(i, 0)
	duration, _ := strconv.ParseInt(date[1], 10, 64)
	Bytes, _ := strconv.ParseInt(date[4], 10, 64)
	r.Time = sql.NullTime{tm.UTC(), true}
	r.Elapsed = sql.NullInt64{duration, true}
	r.Remotehost = sql.NullString{date[2], true}
	r.CodeStatus = sql.NullString{date[3], true}
	r.Bytes = sql.NullInt64{Bytes, true}
	r.Method = sql.NullString{date[5], true}
	r.URL = sql.NullString{date[6], true}
	r.Rfc931 = sql.NullString{date[7], true}
	r.PeerStatus = sql.NullString{date[8], true}
	r.Type = sql.NullString{date[9], true}

	DB.Create(&r)
	return true
}
