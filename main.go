package main

import (
	"log"

	"github.com/gin-gonic/gin"
	couponservice "github.com/ujjawal0619/cm/couponService"
	couponhandler "github.com/ujjawal0619/cm/couponService/handler"
	database "github.com/ujjawal0619/cm/database/couponDB"
	"github.com/ujjawal0619/cm/routes"
)

func main() {
	store, err := database.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	couponService := couponservice.InitCityService(store)
	couponHandler := couponhandler.InitHandler(couponService)

	r := gin.New()

	routes.CouponRoutes(r, couponHandler)

	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	r.Run()
}
