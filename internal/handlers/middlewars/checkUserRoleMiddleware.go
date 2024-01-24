package middlewars

import "github.com/gin-gonic/gin"

func CheckUserRole(role string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userRole, exists := ctx.Get("userRole")
		if !exists {
			ctx.JSON(403, gin.H{"error": "Access forbidden"})
			ctx.Abort()
			return
		}

		if userRole != role {
			ctx.JSON(403, gin.H{"error": "insufficient permissions"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
