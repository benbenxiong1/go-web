package migrations

import (
	"database/sql"
	"go-web/app/model"
	"go-web/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type User struct {
		model.BaseModel
	}

	type Category struct {
		model.BaseModel
	}

	type Topic struct {
		model.BaseModel

		Title      string `gorm:"type:varchar(255);not null;index"`
		Body       string `gorm:"type:longtext;not null"`
		UserId     string `gorm:"type:longtext;not null;index"`
		CategoryId string `gorm:"type:longtext;not null;index"`

		// 会创建 user_id 和 category_id 外键的约束
		User     User
		Category Category

		model.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Topic{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Topic{})
	}

	migrate.Add("2023_07_05_170248_add_topics_table", up, down)
}
