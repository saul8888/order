package config

import "os"

// appConfig type
type appConfig struct {
	ServerPort   string
	DatabaseURI  string
	DatabaseName string
}

// AppConfig stores specific config
var AppConfig = appConfig{}

const (
	defaultServerPort  = "8080"
	defaultDatabaseURI = "mongodb+srv://saul:1234@cluster0-ooeaq.mongodb.net/test?retryWrites=true&w=majority" //"localhost:27017"
	defaultDbName      = "order"                                                                               //"defaultDB"
)

func init() {
	if appPort := os.Getenv("PORT"); len(appPort) == 0 {
		AppConfig.ServerPort = defaultServerPort
	} else {
		AppConfig.ServerPort = appPort
	}

	if databaseURI := os.Getenv("DATABASE_URI"); len(databaseURI) == 0 {
		AppConfig.DatabaseURI = defaultDatabaseURI
	} else {
		AppConfig.DatabaseURI = databaseURI
	}

	if databaseName := os.Getenv("DATABASE_NAME"); len(databaseName) == 0 {
		AppConfig.DatabaseName = defaultDbName
	} else {
		AppConfig.DatabaseName = databaseName
	}

	//fmt.Println(AppConfig)
}
