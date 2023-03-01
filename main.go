package main

import (
	productcontroller "fahmi/controllers"
	"fahmi/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.Connectdb()

	r.GET("/api/produk", productcontroller.Index)
	r.GET("/api/produk/:id", productcontroller.Show)
	r.POST("/api/produk", productcontroller.Create)
	r.PUT("/api/produk/:id", productcontroller.Update)
	r.DELETE("/api/produk/", productcontroller.Delete)

	r.Run()
}
