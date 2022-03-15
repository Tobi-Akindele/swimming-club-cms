package main

import (
	"github.com/joho/godotenv"
	"swimming-club-cms-be/repositories/dbInits"
	"swimming-club-cms-be/routers"
	"swimming-club-cms-be/utils"
)

func main() {

	// Load .env file from specified path. Currently, in the root dir
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	dbInits.InitializeDB()

	router := routers.HandleRequests()
	port := utils.GetEnv(utils.PORT, ":8080")
	err = router.Run(port)
	if err != nil {
		panic(err)
	}
}
