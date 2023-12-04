package noteModel

import "ChatServe/src/utils"

func CreateNote(note *Note) error {
	return utils.SqlCreate(note)
}
func DeleteNote(id string) error {
	return utils.Delete(id, &Note{})
}

func CreateNoteBlackUser(note *BlackList) error {
	return utils.SqlCreate(note)
}
func DeleteNoteBlackUser(id string) error {
	return utils.Delete(id, &BlackList{})
}
