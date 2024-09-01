package main

import (
	"felagi/config"
	"felagi/config/db"
	"felagi/delivery/routers"
)

func main() {
    config.InitiEnvConfigs() 
    db.ConnectDB(config.EnvConfigs.MongoURI)

    router := routers.SetupRouter()

    router.Run(config.EnvConfigs.LocalServerPort)
}


