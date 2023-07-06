package category

import (
	"github.com/gin-gonic/gin"
	"go-web/pkg/app"
	"go-web/pkg/cache"
	"go-web/pkg/database"
	"go-web/pkg/helpers"
	"go-web/pkg/paginator"
	"time"
)

func Get(idstr string) (category Category) {
	database.DB.Where("id", idstr).First(&category)
	return
}

func GetBy(field, value string) (category Category) {
	database.DB.Where("? = ?", field, value).First(&category)
	return
}

func All() (categories []Category) {
	database.DB.Find(&categories)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Category{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (category []Category, paging paginator.Paging) {
	paging = paginator.Paginate(c, database.DB.Model(&Category{}), &category, app.V1Url(database.TableName(&Category{})), perPage)

	return
}

func AllCached() (category []Category) {
	// 设置缓存 key
	cacheKey := "links:all"
	// 设置过期时间
	expireTime := 120 * time.Minute
	// 取数据
	cache.GetObject(cacheKey, &category)

	// 如果数据为空
	if helpers.Empty(category) {
		// 查询数据库
		category = All()
		if helpers.Empty(category) {
			return category
		}
		// 设置缓存
		cache.Set(cacheKey, category, expireTime)
	}
	return
}
