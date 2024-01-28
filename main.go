package main

func main() {
	defer database.closeDB()
	LoadEnv("config.env")
	db := database.InitDB()
	database.connect()

}
