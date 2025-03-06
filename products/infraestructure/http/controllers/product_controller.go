package controllers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"

    "ModaVane/products/application"
    "ModaVane/products/domain"
    "ModaVane/products/infraestructure/http/responses"
)

type ProductController struct {
    createProductUseCase *application.CreateProductUseCase
    getProductUseCase    *application.GetProductUseCase
    updateProductUseCase *application.UpdateProductUseCase
    deleteProductUseCase *application.DeleteProductUseCase
}

func NewProductController(createUC *application.CreateProductUseCase, getUC *application.GetProductUseCase, updateUC *application.UpdateProductUseCase, deleteUC *application.DeleteProductUseCase) *ProductController {
    return &ProductController{
        createProductUseCase: createUC,
        getProductUseCase:    getUC,
        updateProductUseCase: updateUC,
        deleteProductUseCase: deleteUC,
    }
}

func (ctrl *ProductController) Create(ctx *gin.Context) {
    var product domain.Product
    if err := ctx.ShouldBindJSON(&product); err != nil {
        ctx.JSON(http.StatusBadRequest, responses.ErrorResponse("los datos son inválidos", err.Error()))
        return
    }

    idProduct, err := ctrl.createProductUseCase.Execute(product)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse("Error al crear producto", err.Error()))
        return
    }

    product.ID = idProduct
    ctx.JSON(http.StatusCreated, responses.SuccessResponse("Producto creado exitosamente", product))
}

func (ctrl *ProductController) GetAll(ctx *gin.Context) {
    products, err := ctrl.getProductUseCase.ExecuteAll()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse("Error al obtener productos", err.Error()))
        return
    }

    ctx.JSON(http.StatusOK, responses.SuccessResponse("Productos obtenidos exitosamente", products))
}

func (ctrl *ProductController) GetByID(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, responses.ErrorResponse("ID inválido", err.Error()))
        return
    }

    product, err := ctrl.getProductUseCase.ExecuteByID(id)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse("Error al obtener producto", err.Error()))
        return
    }

    ctx.JSON(http.StatusOK, responses.SuccessResponse("Producto obtenido exitosamente", product))
}

func (ctrl *ProductController) Update(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, responses.ErrorResponse("ID inválido", err.Error()))
        return
    }

    var product domain.Product
    if err := ctx.ShouldBindJSON(&product); err != nil {
        ctx.JSON(http.StatusBadRequest, responses.ErrorResponse("Datos inválidos", err.Error()))
        return
    }

    product.ID = id
    if err := ctrl.updateProductUseCase.Execute(product); err != nil {
        ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse("Error al actualizar producto", err.Error()))
        return
    }

    ctx.JSON(http.StatusOK, responses.SuccessResponse("Producto actualizado exitosamente", product))
}

func (ctrl *ProductController) Delete(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, responses.ErrorResponse("ID inválido", err.Error()))
        return
    }

    if err := ctrl.deleteProductUseCase.Execute(id); err != nil {
        ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse("Error al eliminar producto", err.Error()))
        return
    }

    ctx.JSON(http.StatusOK, responses.SuccessResponse("Producto eliminado exitosamente", nil))
}
