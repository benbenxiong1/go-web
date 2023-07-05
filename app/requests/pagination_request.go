package requests

import "github.com/thedevsaddam/govalidator"

type PaginationRequest struct {
	Sort    string `valid:"sort" form:"sort"`
	Order   string `valid:"order" form:"order"`
	PerPage string `valid:"per_page" form:"per_page"`
}

func Pagination(data interface{}) map[string][]string {
	rules := govalidator.MapData{
		"sort":     []string{"in:id,create_at,update_at"},
		"order":    []string{"in:asc,desc"},
		"per_page": []string{"numeric_between:2,100"},
	}

	message := govalidator.MapData{
		"sort": []string{
			"in:排序字段仅支持 id,create_at,update_at",
		},
		"order": []string{
			"in:排序规则仅支持正序（asc）、倒序（desc）",
		},
		"per_page": []string{
			"numeric_between:每页条数的值介于 2-100之间",
		},
	}

	return validate(data, rules, message)
}
