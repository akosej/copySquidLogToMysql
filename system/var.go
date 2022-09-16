package system

var (
	// -- MYSQL

	mysqlHost        = Config("mysql.ip")
	mysqlPort        = Config("mysql.port")
	mysqlUser        = Config("mysql.user")
	mysqlPassword    = Config("mysql.pass")
	mysqlDB          = Config("mysql.db")
	MysqlCredentials = mysqlUser + ":" + mysqlPassword + "@tcp(" + mysqlHost + ":" + mysqlPort + ")/" + mysqlDB + "?charset=utf8&parseTime=True&loc=Local"

	Accesslog = Config("path.AccessLog")
)
