package utils

func SqlCreate(data any) error {
	return DB.Create(data).Error
}
