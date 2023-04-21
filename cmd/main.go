package main

import (
	"cleancodego/internal/apis/controllers"
	application "cleancodego/internal/application/services"
	mysqldb "cleancodego/internal/infrastructure/mysqlDb"
)

func main() {

	//repository depenedecny

	repo := mysqldb.NewMysqlRepository()

	// service dependnecy

	service := application.NewUserService(repo)

	// controller dependency

	controller := controllers.NewUserController(service)

	controller.SaveData()

}
