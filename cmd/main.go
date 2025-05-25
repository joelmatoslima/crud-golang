package main

import (
	controller "lsport/controler"
	"lsport/db"
	"lsport/repository"
	"lsport/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	connection, err := db.ConnectDB()

	if err != nil {
		panic(err)
	}

	var ProductRepository = repository.NewProductRepository(connection)

	var ProductUseCase = usecase.NewProductUseCase(ProductRepository)
	var ProductController = controller.NewProductController(ProductUseCase)

	var server = gin.Default()
	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})

	})

	server.GET("/products", ProductController.GetProducts)
	server.POST("/product", ProductController.CreateProduct)
	server.GET("/product/:id", ProductController.GetProductById)

	server.Run(":8080")

}
