package ping

import (
	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 1000,
		"res":  "success",
	})
}
