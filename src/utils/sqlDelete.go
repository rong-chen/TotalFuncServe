package utils

func Delete(Id string, model interface{}) error {
	return DB.Where("id=?", Id).Delete(model).Error
}
