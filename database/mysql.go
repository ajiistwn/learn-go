package database

var connection string

func init() {
	connection = "MySQL Database Connection"
}

func GetConnection() string {
	return connection
}