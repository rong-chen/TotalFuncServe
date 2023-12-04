package app

import (
	"ChatServe/src/serve/chat"
	friendsServe "ChatServe/src/serve/friends"
	"ChatServe/src/serve/login"
	"ChatServe/src/serve/note"
	"ChatServe/src/serve/ping"
	"ChatServe/src/serve/user"
	"ChatServe/src/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func APP() *gin.Engine {
	r := gin.Default()
	r.Use(Cors())
	r.GET("/connect", chat.Connect)

	r.GET("/ping", ping.Test)
	r.POST("/createUser", user.CreateUser)
	r.POST("/login", login.Login)
	r.GET("/getUserInfo", utils.ValidRequestToken, user.GetUserInfo)
	r.POST("/updateUserInfo", user.SetUserInfo)
	r.POST("/addNote", utils.ValidRequestToken, note.AddNotes)
	r.POST("/setNoteBlackList", utils.ValidRequestToken, note.SetNoteBlackList)
	r.POST("/deleteNote", utils.ValidRequestToken, note.DeleteNote)
	r.POST("/deleteBlackUser", utils.ValidRequestToken, note.DeleteBlackUser)
	r.POST("/GetUserInfoByPhone", utils.ValidRequestToken, user.GetUserInfoByPhone)
	r.POST("/addFriends", utils.ValidRequestToken, friendsServe.AddFriends)
	r.POST("/updateFriends", utils.ValidRequestToken, friendsServe.UpdateFrinds)

	//r.GET("/getRouter", utils.ValidRequestToken, router.GetRouter)
	return r
}
func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token,Authorization,Token,x-token,Sec-WebSocket-Protocol")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PATCH, PUT")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type, ")
		context.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
	}
}
