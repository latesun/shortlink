package models

import (
	"time"

	"github.com/imiskolee/optional"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type GWShortener struct {
	CreatedAt  time.Time       `gorm:"column:created_at; type:datetime ;" json:"created_at"`
	ExpireDay  optional.Int    `gorm:"column:expire_day; type:int(11) ;default:'365'" json:"expire_day"`
	ExpiredAt  *time.Time      `gorm:"column:expired_at; type:datetime ;" json:"expired_at"`
	ID         optional.Int    `gorm:"column:id; type:int(11) AUTO_INCREMENT;" json:"id"`
	LongURL    optional.String `gorm:"column:long_url; type:varchar(255) ;" json:"long_url"`
	ShortenKey optional.String `gorm:"column:shorten_key; type:varchar(255) ;" json:"shorten_key"`
	SourceID   optional.Int    `gorm:"column:source_id; type:int(11) ;" json:"source_id"`
	SourceType optional.String `gorm:"column:source_type; type:varchar(191) ;" json:"source_type"`
	StoreID    optional.Int    `gorm:"column:store_id; type:int(11) ;" json:"store_id"`
	UpdatedAt  time.Time       `gorm:"column:updated_at; type:datetime ;" json:"updated_at"`
}

func (s *GWShortener) TableName() string {
	return "shorteners"
}
