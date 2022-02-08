package models

import "go-web/pkg/database"

func IsExist(model interface{},field string,value interface{}) bool {
	var count int64
	database.DB.Model(model).Where(field+" = ?", value).Count(&count)
	return count > 0
}
