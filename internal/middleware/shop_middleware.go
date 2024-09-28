package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sigit14ap/warehouse-service/helpers"
	"github.com/sigit14ap/warehouse-service/internal/services"
)

func ShopMiddleware(shopService *services.ShopService) gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("Authorization")
		if token == "" {
			helpers.ErrorResponse(context, http.StatusUnauthorized, "Authorization required")
			context.Abort()
			return
		}

		shopDetail, err := shopService.ShopDetail(token)
		if err != nil {
			helpers.ErrorResponse(context, http.StatusUnauthorized, "Shop does not allowed")
			return
		}

		context.Set("shopID", shopDetail.ID)
		context.Next()
	}
}
