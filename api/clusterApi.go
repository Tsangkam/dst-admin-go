package api

import (
	"dst-admin-go/config/global"
	"dst-admin-go/constant/consts"
	"dst-admin-go/model"
	"dst-admin-go/service"
	"dst-admin-go/vo"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ClusterApi struct{}

var clusterManager = service.ClusterManager{}

// var clusterService = service.HomeService{}

//func (c *ClusterApi) GetGameConfig(ctx *gin.Context) {
//	ctx.JSON(http.StatusOK, vo.Response{
//		Code: 200,
//		Msg:  "success",
//		Data: clusterService.GetGameConfig(ctx),
//	})
//}
//
//func (c *ClusterApi) SaveGameConfig(ctx *gin.Context) {
//
//	gameConfig := level.GameConfig{}
//	ctx.ShouldBind(&gameConfig)
//	fmt.Printf("%v", gameConfig.Caves.ServerIni)
//	clusterService.SaveGameConfig(ctx, &gameConfig)
//
//	ctx.JSON(http.StatusOK, vo.Response{
//		Code: 200,
//		Msg:  "success",
//		Data: nil,
//	})
//}

func (c *ClusterApi) GetClusterList(ctx *gin.Context) {
	clusterManager.QueryCluster(ctx)
}

func (c *ClusterApi) CreateCluster(ctx *gin.Context) {

	clusterModel := model.Cluster{}
	err := ctx.ShouldBind(&clusterModel)
	if err != nil {
		log.Panicln("参数错误")
	}
	log.Println(clusterModel)

	if clusterModel.SteamCmd == "" || clusterModel.ClusterName == "" {
		log.Panicln("参数错误, steamcmd 或者 clusterName 不能为空")
	}
	if clusterModel.ForceInstallDir == "" {
		clusterModel.ForceInstallDir = filepath.Join(consts.HomePath, "dst-dedicated-cluster", clusterModel.ClusterName)
	}
	if clusterModel.Backup == "" {
		clusterModel.Backup = consts.KleiDstPath
	}
	if clusterModel.ModDownloadPath == "" {
		clusterModel.ModDownloadPath = consts.KleiDstPath
	}

	clusterManager.CreateCluster(&clusterModel)
	global.CollectMap.AddNewCollect(clusterModel.ClusterName)

	ctx.JSON(http.StatusOK, vo.Response{
		Code: 200,
		Msg:  "success",
		Data: nil,
	})

}

func (c *ClusterApi) UpdateCluster(ctx *gin.Context) {
	clusterModel := model.Cluster{}
	ctx.ShouldBind(&clusterModel)
	fmt.Printf("%v", clusterModel)
	clusterManager.UpdateCluster(&clusterModel)

	ctx.JSON(http.StatusOK, vo.Response{
		Code: 200,
		Msg:  "success",
		Data: nil,
	})

}

func (c *ClusterApi) DeleteCluster(ctx *gin.Context) {

	var id int

	if idParam, isExist := ctx.GetQuery("id"); isExist {
		id, _ = strconv.Atoi(idParam)
	}

	clusterModel, err := clusterManager.DeleteCluster(uint(id))
	if err != nil {
		log.Panicln("delete cluster error", err)
	}

	global.CollectMap.RemoveCollect(clusterModel.ClusterName)

	ctx.JSON(http.StatusOK, vo.Response{
		Code: 200,
		Msg:  "success",
		Data: nil,
	})
}
