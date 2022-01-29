package po

import (
	"fmt"

	"gorm.io/gorm"
)

type Instance struct {
	gorm.Model
	UUID  string `gorm:"uuid"`
	Stage string `gorm:"stage"`
}

func (inst *Instance) Create(db *gorm.DB) error {
	result := db.Create(inst)
	if result.Error != nil {
		return fmt.Errorf("create instance failed: %w", result.Error)
	}
	return nil
}
