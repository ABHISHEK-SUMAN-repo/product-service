package setup

import (
	"log"
	"product-service/config"
	"product-service/router"
)

func Initializer(env string) {
	err := config.PostgresConnection(env)
	if err != nil {
		log.Panic("Postgres connection error")
	}
	log.Print("Connected to Postgres")
	router.RoutingInitialize()
}

