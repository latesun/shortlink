package models

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Shortener struct {
	CreatedAt  time.Time  `gorm:"column:created_at; type:datetime"`
	ExpireDay  int        `gorm:"column:expire_day; type:int(11); default:'365'" `
	ExpiredAt  *time.Time `gorm:"column:expired_at; type:datetime"`
	ID         int        `gorm:"column:id; type:int(11) AUTO_INCREMENT"`
	LongURL    string     `gorm:"column:long_url; type:varchar(255)"`
	ShortenKey string     `gorm:"column:shorten_key; type:varchar(255)"`
	SourceID   int        `gorm:"column:source_id; type:int(11)"`
	SourceType string     `gorm:"column:source_type; type:varchar(191)"`
	StoreID    int        `gorm:"column:store_id; type:int(11)"`
	UpdatedAt  time.Time  `gorm:"column:updated_at; type:datetime"`
}

func (s *Shortener) TableName() string {
	return "shorteners"
}
