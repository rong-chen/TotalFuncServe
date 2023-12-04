package utils

import "github.com/gin-gonic/gin"

func BackSuccessResp() gin.H {
	return gin.H{
		"code": 200,
		"msg":  "success",
	}
}

func BackFailedResp() gin.H {
	return gin.H{
		"code": 201,
		"msg":  "failed",
	}
}
func BackMessageResp(code int, message string) gin.H {
	return gin.H{
		"code": code,
		"msg":  message,
	}
}
func BackDataResp(code int, message string, val interface{}) gin.H {
	return gin.H{
		"code": code,
		"msg":  message,
		"data": val,
	}
}
