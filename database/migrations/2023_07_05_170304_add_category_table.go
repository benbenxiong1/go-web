package migrations

import (
	"database/sql"
	"go-web/app/model"
	"go-web/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type Category struct {
		model.BaseModel

		Name        string `gorm:"type:varchar(255);not null;index"`
		Description string `gorm:"type:varchar(255);default:null"`

		model.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Category{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Category{})
	}

	migrate.Add("2023_07_05_170304_add_category_table", up, down)
}
