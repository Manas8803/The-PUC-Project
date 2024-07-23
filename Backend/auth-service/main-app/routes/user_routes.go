package routes

import (
	"github.com/Manas8803/The-PUC-Project__BackEnd/auth-service/main-app/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.RouterGroup) {
	router.POST("/auth/register", controllers.Register)
	router.POST("/auth/login", controllers.Login)
}
