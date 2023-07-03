package migrate

import (
	"go-web/pkg/console"
	"go-web/pkg/database"
	"go-web/pkg/file"
	"gorm.io/gorm"
	"os"
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

func (m *Migrator) CreateMigrationsTable() {
	migration := Migration{}

	// 不存在才创建
	if !m.Migrator.HasTable(&migration) {
		err := m.Migrator.CreateTable(&migration)
		if err != nil {
			return
		}
	}
}

func (m *Migrator) Up() {
	// 读取所有迁移文件，确保按照时间排序
	migrationFiles := m.readAllMigrationFiles()

	// 获取此次的批次值
	batch := m.getBatch()

	// 获取所有迁移数据
	var migrations []Migration
	m.DB.Find(&migrations)

	// 可以通过此值判断数据库是否已更新
	runed := false

	// 对迁移文件进行遍历，如果没有执行过，就执行 up 回调
	for _, mfile := range migrationFiles {
		// 对比文件名称，看是否已经执行过
		if mfile.isNotMigrated(migrations) {
			m.runUpMigration(mfile, batch)
			runed = true
		}
	}

	if !runed {
		console.Success("database is up to date.")
	}

}

// readAllMigrationFiles 从文件目录读取文件，确保正确的时间排序
func (m *Migrator) readAllMigrationFiles() []MigrationFile {
	// 读取 database/migration/ 目录下的所有文件
	files, err := os.ReadDir(m.Folder)
	console.ExitIf(err)

	var migrateFiels []MigrationFile
	for _, f := range files {
		// 去除文件后缀 .go
		fileName := file.FileNameWithoutExtension(f.Name())

		// 通过迁移文件的名称获取 MigrationFile 对象
		mfile := getMigrationFile(fileName)

		// 判断迁移文件是否可用
		if len(mfile.FileName) > 0 {
			migrationFiles = append(migrateFiels, mfile)
		}
	}

	// 返回排序好的 MigrationFile 数组
	return migrateFiels
}

// getBatch 获取当前批次值
func (m *Migrator) getBatch() int {
	// 默认为 1
	batch := 1

	// 取最后一条迁移记录
	lastMigration := Migration{}
	m.DB.Order("id DESC").First(&lastMigration)

	// 如果有值 +1
	if lastMigration.ID > 0 {
		batch = lastMigration.batch + 1
	}

	return batch
}

// runUpMigration 执行迁移（执行迁移的up方法）
func (m *Migrator) runUpMigration(mfile MigrationFile, batch int) {
	// 执行 up 区域快的 sql
	if mfile.Up != nil {
		// 友好提示
		console.Warning("migrating " + mfile.FileName)
		// 执行 up 方法
		mfile.Up(database.DB.Migrator(), database.SQLDB)
		// 提示已迁移了哪个文件
		console.Success("migrated " + mfile.FileName)
	}

	// 入库
	err := m.DB.Create(&Migration{
		Migration: mfile.FileName,
		batch:     batch,
	}).Error

	console.ExitIf(err)
}
