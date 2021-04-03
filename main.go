package main

import (
	adapter "user-api/adapter/db"
	adapterHttp "user-api/adapter/http"
	"user-api/external/http"

	"user-api/external/db"
	"user-api/external/db/test/helper"
	"user-api/usecase"
)

func main() {

	helper := helper.NewMongoDBHelper()

	err := helper.StartMongoDB()
	nosqlDB, err := db.NewNoSQLDB(helper.DatabaseURI(), helper.DatabaseName())
	userRepository, err := adapter.NewUserRepository(&nosqlDB)
	usecase, err := usecase.NewCreateUser(&userRepository)
	controller, err := adapterHttp.NewHttpController(&usecase)
	handler, err := http.NewHandler(&controller)
	register := http.RegisterHandlers(&handler)

	if err != nil {
		panic(err)
	}

	register.Run()
}
