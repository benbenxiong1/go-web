//Package topic 模型
package topic

import (
	"go-web/app/model"
	"go-web/app/model/category"
	"go-web/app/model/user"
	"go-web/pkg/database"
)

type Topic struct {
	model.BaseModel

	// Put fields in here
	Title      string `json:"title,omitempty" `
	Body       string `json:"body,omitempty" `
	UserID     string `json:"user_id,omitempty"`
	CategoryID string `json:"category_id,omitempty"`

	// 通过user_id 关联用户
	User user.User `json:"user"`

	// 通过category_id 去关联分类
	Category category.Category `json:"category"`

	model.CommonTimestampsField
}

func (topic *Topic) Create() {
	database.DB.Create(&topic)
}

func (topic *Topic) Save() (rowsAffected int64) {
	result := database.DB.Save(&topic)
	return result.RowsAffected
}

func (topic *Topic) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&topic)
	return result.RowsAffected
}
