package config

const (
	//下面2个参数要从环境变量中读取
	ORDER_SERVER_HOST = "ORDER_SERVER_HOST"
	USER_SERVER_HOST  = "USER_SERVER_HOST"

	ORDER_SERVER_HOST_DEFAULT = "localhost:13332"
	ORDER_SERVER_PATH         = "gapporder"

	USER_SERVER_HOST_DEFAULT = "localhost:13331"
	USER_SERVER_PATH         = "gappuser"
)
