package seeders

import "go-web/pkg/seed"

func Initialize() {

	// 触发本目录下的其他 init 方法

	// 指定优先于目录下的其他文件运行
	seed.SetRunOrder([]string{
		"SeedUsersTable",
	})
}
