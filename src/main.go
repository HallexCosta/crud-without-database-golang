package main

import "crud-without-database-golang/src/config"

func init() {
	_AppConfig := config.NewAppConfig()
	_Persistence := _AppConfig.NewPersistence("users")

	_Persistence.InitPersistence()
}

func main() {
	App()
	Routes()

	AppListen(8080)
}
