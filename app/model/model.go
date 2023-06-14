// Package model  user : benbenxiong  time : 2023-07-2023/6/7 20:20:19
package model

import (
	"github.com/spf13/cast"
	"time"
)

// BaseModel 模型基类
type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`
}

// CommonTimestampsField 时间戳
type CommonTimestampsField struct {
	CreatedAt time.Time `gorm:"column:created_at;index;" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;index;" json:"updated_at,omitempty"`
}

func (b *BaseModel) GetStringId() string {
	return cast.ToString(b.ID)
}
