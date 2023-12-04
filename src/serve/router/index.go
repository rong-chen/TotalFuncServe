package router

import (
	routers "ChatServe/src/model/router"
	"ChatServe/src/utils"
	"github.com/gin-gonic/gin"
)

func GetRouter(c *gin.Context) {
	routerList := []routers.Router{}
	router := routers.GetRouters()
	for i := 0; i < len(router); i++ {
		routerList = append(routerList, router[i])
	}
	c.JSON(200, utils.BackDataResp(200, "查询成功", routerList))
}
