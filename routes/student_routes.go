package routes

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"go-web-mini/controller"
	"go-web-mini/middleware"
)

// 注册用户路由
func InitStudentRoutes(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
	studentController := controller.NewStudentController()
	router := r.Group("/student")
	// 开启jwt认证中间件
	router.Use(authMiddleware.MiddlewareFunc())
	// 开启casbin鉴权中间件
	router.Use(middleware.CasbinMiddleware())
	{
		router.GET("/list", studentController.GetStudents)
		router.POST("/create", studentController.CreateStudent)
		router.PATCH("/update", studentController.UpdateStudent)
		router.DELETE("/inClass/batch", studentController.BatchInClass)
		router.DELETE("/delete/batch", studentController.BatchDeleteStudentByIds)
	}
	return r
}
