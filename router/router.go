package router

import (
	"log"
	"product-service/controller"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// RoutingInitialize initializes the routes
func RoutingInitialize() {
	r := gin.Default()
	basePath := viper.GetString("base_path")
	port := viper.GetString("port")

	routeGroup := r.Group(basePath)
	log.Printf("route group: %v", routeGroup.BasePath())
	productRouter(routeGroup)
	
	address := ":" + port
	log.Print("Running server on address", address)
	if err := r.Run(address); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}

func productRouter(router *gin.RouterGroup) {
	// router.GET("/test", controller.Test)
	router.POST("/create/products", controller.CreateProducts)
	router.GET("/products/id",controller.GetProductsByIds)

}