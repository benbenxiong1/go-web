package database

import (
	"database/sql"
	"errors"
	"fmt"
	"go-web/pkg/config"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var SQLDB *sql.DB

func Connect(dbConfig gorm.Dialector, _logger logger.Interface) {
	// 使用 gorm.Open 连接数据库
	var err error
	DB, err = gorm.Open(dbConfig, &gorm.Config{
		Logger: _logger,
	})

	// 处理错误
	if err != nil {
		fmt.Println(err.Error())
	}

	// 获取底层的 sqlDB
	SQLDB, err = DB.DB()
	if err != nil {
		fmt.Println(err.Error())
	}

}

// CurrentDatabase 获取当前数据库名称
func CurrentDatabase() (dbName string) {
	dbName = DB.Migrator().CurrentDatabase()
	return
}

// DeleteAllTable 删除所有数据表
func DeleteAllTable() error {
	var err error
	switch config.Get("database.connection") {
	case "mysql":
		err = deleteMysqlTables()
	case "sqlite":
		err = deleteAllSqliteTables()
	default:
		panic(errors.New("database connection not supported"))
	}
	return err
}

func deleteAllSqliteTables() error {
	var tables []string

	// 读取所有数据表
	err := DB.Select(&tables, "SELECT name FROM sqlite_master WHERE type = `table`").Error
	if err != nil {
		return err
	}

	// 删除所有表
	for _, table := range tables {
		err := DB.Migrator().DropTable(table)
		if err != nil {
			return err
		}
	}

	return nil
}

func deleteMysqlTables() error {
	dbName := CurrentDatabase()
	var tables []string

	// 读取所有数据表
	err := DB.Table("information_schema.tables").Where("table_schema = ?", dbName).
		Pluck("table_name", &tables).Error
	if err != nil {
		return err
	}

	// 暂时关闭外键检测
	DB.Exec("SET foreign_key_checks = 0;")

	// 删除所有表
	for _, table := range tables {
		err := DB.Migrator().DropTable(table)
		if err != nil {
			return err
		}
	}

	DB.Exec("SET foreign_key_checks = 1;")

	return nil
}

func TableName(obj interface{}) string {
	stmt := &gorm.Statement{DB: DB}
	err := stmt.Parse(obj)
	if err != nil {
		return ""
	}
	return stmt.Schema.Table
}
