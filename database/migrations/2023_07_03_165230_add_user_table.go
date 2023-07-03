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

		Name     string `gorm:"type:varchar(255);not null;index"`
		Email    string `gorm:"type:varchar(255);index;default:null"`
		Phone    string `gorm:"type:varchar(20);index;default:null"`
		Password string `gorm:"type:varchar(255)"`

		model.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&User{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&User{})
	}

	migrate.Add("2023_07_03_165230_add_user_table", up, down)
}
