package note

import (
	noteModel "ChatServe/src/model/note"
	"ChatServe/src/utils"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// AddNotes 添加日记
func AddNotes(c *gin.Context) {
	var data noteModel.Note
	c.Bind(&data)
	data.Id = uuid.NewV4()
	err := noteModel.CreateNote(&data)
	if err != nil {
		panic(err)
	}
	c.JSON(200, utils.BackSuccessResp())
}

// DeleteNote 删除日记
func DeleteNote(c *gin.Context) {
	type form struct {
		Id string
	}
	var data form
	c.Bind(&data)
	err := noteModel.DeleteNote(data.Id)
	if err != nil {
		c.JSON(200, utils.BackFailedResp())
		return
	}
	c.JSON(200, utils.BackSuccessResp())
}

// SetNoteBlackList 设置黑名单
func SetNoteBlackList(c *gin.Context) {
	data := new(noteModel.BlackList)
	c.Bind(data)
	data.Id = uuid.NewV4()
	err := noteModel.CreateNoteBlackUser(data)
	if err != nil {
		c.JSON(200, utils.BackMessageResp(201, "该角色或该日记不存在"))
		return
	}
	c.JSON(200, utils.BackSuccessResp())
}

// DeleteBlackUser 删除黑名单角色
func DeleteBlackUser(c *gin.Context) {
	type form struct {
		Id string
	}
	var data form
	c.Bind(&data)
	err := noteModel.DeleteNoteBlackUser(data.Id)
	if err != nil {
		c.JSON(200, utils.BackFailedResp())
		return
	}
	c.JSON(200, utils.BackSuccessResp())
}
