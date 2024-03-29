package topic

import (
	"github.com/gin-gonic/gin"
	"go-web/pkg/app"
	"go-web/pkg/database"
	"go-web/pkg/paginator"
	"gorm.io/gorm/clause"
)

func Get(idstr string) (topic Topic) {
	database.DB.Preload(clause.Associations).Where("id", idstr).First(&topic)
	return
}

func GetBy(field, value string) (topic Topic) {
	database.DB.Where("? = ?", field, value).First(&topic)
	return
}

func All() (topics []Topic) {
	database.DB.Find(&topics)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Topic{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (topic []Topic, paging paginator.Paging) {
	paging = paginator.Paginate(c, database.DB.Model(&Topic{}), &topic, app.V1Url(database.TableName(&Topic{})), perPage)

	return
}
