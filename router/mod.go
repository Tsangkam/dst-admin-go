package router

import (
	"dst-admin-go/api"
	"github.com/gin-gonic/gin"
)

func initModRouter(router *gin.RouterGroup) {

	modApi := api.ModApi{}
	mod := router.Group("/api/mod")
	{
		mod.GET("/search", modApi.SearchModList)
		mod.GET("/:modId", modApi.GetModInfo)
		mod.PUT("/:modId", modApi.UpdateMod)
		mod.GET("", modApi.GetMyModList)
		mod.DELETE("/:modId", modApi.DeleteMod)
		mod.DELETE("/setup/workshop", modApi.DeleteSetupWorkshop)
		mod.GET("/modinfo/:modId", modApi.GetModInfoFile)
		mod.POST("/modinfo", modApi.SaveModInfoFile)
	}

}
