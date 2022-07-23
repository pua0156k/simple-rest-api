package router

import (
	"simpleRestAPI/controller"

	"github.com/gin-gonic/gin"
)

/*
InnitRouter 路由初始化
*/
func InitRouter(router *gin.Engine) {
	quiz := router.Group("/quiz/v1")
	{
		quiz.POST("/comment", controller.AddComment)
		quiz.GET("/comment/:id", controller.GetComment)
		quiz.PUT("/comment/:id", controller.UpdateComment)
		quiz.DELETE("/comment/:id", controller.DeleteComment)
	}
	/*machine := router.Group("/equip")
	{
		machine.GET("/query", controller.GetMachineBarBySize)
		machine.GET("/query/all", controller.GetMachineTotalBySize)
		machine.GET("/querydetail", controller.GetMachineDetailBySize)
		machine.GET("/query/mach", controller.GetMachineById)
		machine.GET("/query/bar/mach", controller.GetMachineBarById)
		machine.GET("/query/timeline/mach", controller.GetTimelineById)
		machine.GET("/query/machlist", controller.GetMachId)
	}*/

}
