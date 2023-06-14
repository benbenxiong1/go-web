package user

import (
	"go-web/pkg/hash"
	"gorm.io/gorm"
)

// BeforeSave GORM 的模型钩子，在创建和更新模型前调用
func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	if !hash.BcryptIsHashed(u.Password) {
		u.Password = hash.BcryptHash(u.Password)
	}
	return
}
