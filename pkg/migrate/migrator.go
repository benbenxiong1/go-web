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
	Batch     int
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

// Up 执行未迁移过的文件
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

// Rollback 回滚上一个操作
func (m *Migrator) Rollback() {
	// 获取最后一批次的迁移数据
	lastMigration := Migration{}
	m.DB.Order("id DESC").First(&lastMigration)
	var migrations []Migration
	m.DB.Where("batch = ?", lastMigration.Batch).Order("id DESC").Find(&migrations)

	// 回滚最后一批次迁移的数据
	if !m.rollbackMigrations(migrations) {
		console.Success("[migrations] table is empty, nothing to rollback.")
	}

}

// Reset 回滚所有迁移
func (m *Migrator) Reset() {
	var migrations []Migration

	// 按照顺序读取所有迁移文件
	m.DB.Order("id DESC").Find(&migrations)

	// 回滚所有迁移
	if m.rollbackMigrations(migrations) {
		console.Success("[migrations] table is empty, nothing to rollback.")
	}
}

// Refresh 回滚所有迁移，并再次执行所有迁移
func (m *Migrator) Refresh() {
	// 回滚所有迁移
	m.Reset()

	// 再次执行所有迁移
	m.Up()
}

func (m *Migrator) Fresh() {
	// 获取数据库名称，用以提示
	dbName := database.CurrentDatabase()

	// 删除所有表
	err := database.DeleteAllTable()
	console.ExitIf(err)

	console.Success("clearup database " + dbName)

	// 重新创建 migrate 表
	m.CreateMigrationsTable()
	console.Success("[migrations] table created.")

	// 重新调用 up 命令
	m.Up()
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
			migrateFiels = append(migrateFiels, mfile)
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
		batch = lastMigration.Batch + 1
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
		Batch:     batch,
	}).Error

	console.ExitIf(err)
}

// rollbackMigrations 回退迁移 按照顺序执行迁移文件的 down 方法
func (m *Migrator) rollbackMigrations(migrations []Migration) bool {
	// 迁移回退的标志
	runed := false

	for _, _migration := range migrations {
		// 友好提示
		console.Warning("rollback " + _migration.Migration)

		// 执行迁移文件的 down 方法
		mfile := getMigrationFile(_migration.Migration)
		if mfile.Down != nil {
			mfile.Down(database.DB.Migrator(), database.SQLDB)
		}

		runed = true

		// 回退成功了 删除这条记录
		m.DB.Delete(&_migration)

		// 打印运行状态
		console.Success("finish " + mfile.FileName)
	}

	return runed
}
