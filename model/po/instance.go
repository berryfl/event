package po

import (
	"gorm.io/gorm"
)

type Instance struct {
	gorm.Model
	UUID  string `gorm:"uuid"`
	Stage string `gorm:"stage"`
}
