package policies

import (
	"github.com/gin-gonic/gin"
	"go-web/app/model/topic"
	"go-web/pkg/auth"
)

func CanModifyTopic(c *gin.Context, _topic topic.Topic) bool {
	return auth.CurrentUid(c) == _topic.UserID
}
