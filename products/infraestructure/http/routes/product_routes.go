package routes

import (
    "github.com/gin-gonic/gin"
    "ModaVane/products/infraestructure/http/controllers"
)

func SetupProductRoutes(router *gin.Engine, controller *controllers.ProductController) {
    productRoutes := router.Group("/products")
    {
        productRoutes.POST("/", controller.Create)
        productRoutes.GET("/", controller.GetAll)
        productRoutes.GET("/:id", controller.GetByID)
        productRoutes.PUT("/:id", controller.Update)
        productRoutes.DELETE("/:id", controller.Delete)
    }
}
