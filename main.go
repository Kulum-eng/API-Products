package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	p_application "ModaVane/products/application"
	p_adapters "ModaVane/products/infraestructure/adapters"
	p_controllers "ModaVane/products/infraestructure/http/controllers"
	p_routes "ModaVane/products/infraestructure/http/routes"
	"ModaVane/core"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("CORS")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	myGin := gin.New()
	myGin.RedirectTrailingSlash = false

	myGin.Use(CORS())

	db, err := core.InitDB()
	if err != nil {
		log.Println(err)
		return
	}

	productRepository := p_adapters.NewMySQLProductRepository(db)
	createProductUseCase := p_application.NewCreateProductUseCase(productRepository)
	getProductUseCase := p_application.NewGetProductUseCase(productRepository)
	updateProductUseCase := p_application.NewUpdateProductUseCase(productRepository)
	deleteProductUseCase := p_application.NewDeleteProductUseCase(productRepository)

	createProductController := p_controllers.NewProductController(createProductUseCase, getProductUseCase, updateProductUseCase, deleteProductUseCase)
	p_routes.SetupProductRoutes(myGin, createProductController)

	myGin.Run(":8080")
}
