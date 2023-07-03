package migrate

import (
	"database/sql"
	"gorm.io/gorm"
)

// migrationFunc 定义 up 和 down 回调方法的类型
type migrationFunc func(migrator gorm.Migrator, db *sql.DB)

// MigrationFile 代表单个迁移文件
type MigrationFile struct {
	Up       migrationFunc
	Down     migrationFunc
	FileName string
}

// migrationFiles 所有迁移文件数据
var migrationFiles []MigrationFile

// Add 新增一个迁移文件，所有的迁移文件都需要调用此方法来注册
func Add(name string, up migrationFunc, down migrationFunc) {
	migrationFiles = append(migrationFiles, MigrationFile{
		Up:       up,
		Down:     down,
		FileName: name,
	})
}