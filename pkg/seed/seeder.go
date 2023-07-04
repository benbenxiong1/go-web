package seed

import (
	"go-web/pkg/console"
	"go-web/pkg/database"
	"gorm.io/gorm"
)

// 存放所有 Seeder
var seeders []Seeder

var orderedSeederNames []string

type SeederFunc func(*gorm.DB)

// Seeder 对应每一个 database/seeders 目录下的 Seeder 文件
type Seeder struct {
	Func SeederFunc
	Name string
}

// Add 注册到 seeders 数组中
func Add(name string, fu SeederFunc) {
	seeders = append(seeders, Seeder{
		Func: fu,
		Name: name,
	})
}

// SetRunOrder 设置按循序执行的 Seeder 数组
func SetRunOrder(name []string) {
	orderedSeederNames = name
}

// GetSeeder 通过名称来获取 seeder 对象
func GetSeeder(name string) Seeder {
	for _, sdr := range seeders {
		if name == sdr.Name {
			return sdr
		}
	}
	return Seeder{}
}

// RunAll 运行所有 Seeder
func RunAll() {
	// 线运行 ordered 的
	executed := make(map[string]string)
	for _, name := range orderedSeederNames {
		sdr := GetSeeder(name)
		if len(sdr.Name) > 0 {
			console.Warning("Running Ordered Seeder: " + sdr.Name)
			sdr.Func(database.DB)
			executed[name] = name
		}
	}

	// 在运行剩下的

	for _, sdr := range seeders {
		// 过滤已运行的
		if _, ok := executed[sdr.Name]; !ok {
			console.Warning("Running Seeder:" + sdr.Name)
			sdr.Func(database.DB)
		}
	}
}

// RunSeeder 运行单个 Seeder
func RunSeeder(name string) {
	for _, sdr := range seeders {
		if name == sdr.Name {
			console.Warning("Running Seeder:" + sdr.Name)
			sdr.Func(database.DB)
			break
		}
	}
}
