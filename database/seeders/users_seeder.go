package seeders

import (
	"fmt"
	"go-web/database/factories"
	"go-web/pkg/console"
	"go-web/pkg/logger"
	"go-web/pkg/seed"
	"gorm.io/gorm"
)

func init() {
	// 添加 Seeder
	seed.Add("SeedUsersTable", func(db *gorm.DB) {

		// 创建10个用户对象
		users := factories.MakeUsers(10)

		// 批量创建用户（批量创建不会触发模型钩子）
		result := db.Table("users").Create(&users)

		// 记录错误
		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		// 打印运行情况
		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))

	})
}
