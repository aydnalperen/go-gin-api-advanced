package books

import (
	"net/http"

	"github.com/aydnalperen/go-gin-api-advanced/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h Handler) GetBooks(ctx *gin.Context) {
	var books []models.Book

	if result := h.DB.Find(&books); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	ctx.JSON(http.StatusOK, &books)
}
