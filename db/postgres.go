package database

import "os"

type dbConfig struct {
	dbHost     string
	dbName     string
	dbPort     string
	dbUserName string
	dbPassword string
}

func InitDB() *dbConfig {
	dbConfig := dbConfig{}
	dbConfig.dbHost = os.Getenv("POSTGRES_HOST")
	dbConfig.dbName = os.Getenv("POSTGRES_DBNAME")
	dbConfig.dbPort = os.Getenv("POSTGRES_PORT")
	dbConfig.dbUserName = os.Getenv("POSTGRES_USER")
	dbConfig.dbPassword = os.Getenv("POSTGRES_PASSWORD")

	return &dbConfig
}
