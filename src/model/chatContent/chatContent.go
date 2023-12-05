package chatContent

import "ChatServe/src/utils"

func CreateChatContent(cc *ChatContent) error {
	return utils.DB.Create(&cc).Error
}
