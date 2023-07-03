package migrate

import (
	"go-web/pkg/database"
	"gorm.io/gorm"
)

// Migrator 数据迁移操作类
type Migrator struct {
	Folder   string
	DB       *gorm.DB
	Migrator gorm.Migrator
}

// Migration 对应数据的 Migrations 表里的一条数据
type Migration struct {
	ID        uint64 `gorm:"primaryKey,autoIncrement;"`
	Migration string `gorm:"type:varchar(255);not null;unique;"`
	batch     int
}

func NewMigrator() *Migrator {
	// 初始化必要属性
	migrator := &Migrator{
		Folder:   "database/migrations/",
		DB:       database.DB,
		Migrator: database.DB.Migrator(),
	}

	// 不存在 则创建迁移记录表
	migrator.CreateMigrationsTable()

	return migrator
}

func (m Migrator) CreateMigrationsTable() {
	migration := Migration{}

	// 不存在才创建
	if !m.Migrator.HasTable(&migration) {
		err := m.Migrator.CreateTable(&migration)
		if err != nil {
			return
		}
	}
}