package system

var (
	// -- MYSQL

	mysqlHost        = Config("mysql.ip")
	mysqlPort        = Config("mysql.port")
	mysqlUser        = Config("mysql.user")
	mysqlPassword    = Config("mysql.pass")
	mysqlDB          = Config("mysql.db")
	MysqlCredentials = mysqlUser + ":" + mysqlPassword + "@tcp(" + mysqlHost + ":" + mysqlPort + ")/" + mysqlDB + "?charset=utf8&parseTime=True&loc=Local"
	ConnectType      = []string{"TCP_CLIENT_REFRESH_MISS/200", "TCP_MISS/200", "TCP_MISS_ABORTED/200", "TCP_TUNNEL/200", "TCP_TUNNEL_ABORTED/200", "TCP_REFRESH_UNMODIFIED/200", "TCP_NEGATIVE_HIT/200", "TCP_REFRESH_MISS/200", "TCP_SWAPFAIL_MISS/200"}
	Accesslog        = Config("path.AccessLog")
)
