package routes

import (
	"github.com/brnocorreia/ufba-enhacer/internal/controller"
	"github.com/gin-gonic/gin"
)

func InitUserRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {
	r.POST("/createUser", userController.CreateUser)
}
