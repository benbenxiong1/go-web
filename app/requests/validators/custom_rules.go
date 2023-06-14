package validators

import (
	"errors"
	"fmt"
	"github.com/thedevsaddam/govalidator"
	"go-web/pkg/database"
	"strings"
)

func init() {
	// 自定义规则 not_exists，验证请求数据必须不存在于数据库中。
	// 常用于保证数据库某个字段的值唯一，如用户名、邮箱、手机号、或者分类的名称。
	// not_exists 参数可以有两种，一种是 2 个参数，一种是 3 个参数：
	// not_exists:users,email 检查数据库表里是否存在同一条信息
	// not_exists:users,email,32 排除用户掉 id 为 32 的用户
	govalidator.AddCustomRule("not_exists", func(field string, rule string, message string, value interface{}) error {
		rng := strings.Split(strings.TrimPrefix(rule, "not_exists:"), ",")

		// 第一个参数 数据表
		tableName := rng[0]
		// 第二个参数 数据表字段
		fieldName := rng[1]

		// 第三个参数 排除当前 ID
		var Id string
		if len(rng) > 2 {
			Id = rng[2]
		}

		// 用户请求过来的数据
		requestValue := value.(string)

		// 组合sql
		query := database.DB.Table(tableName).Where(fieldName+" = ?", requestValue)

		// 如果第三个参数存在 则拼接
		if len(Id) > 0 {
			query.Where("id <> ?", Id)
		}

		// 查询数据库
		var count int64
		query.Count(&count)

		if count != 0 {
			// 如果有自定义错误消息的话
			if message != "" {
				return errors.New(message)
			}
			// 默认的错误消息
			return fmt.Errorf("%v 已被占用", requestValue)
		}
		// 验证通过
		return nil
	})
}