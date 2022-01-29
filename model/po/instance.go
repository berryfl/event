package po

import (
	"fmt"

	"gorm.io/gorm"
)

type Instance struct {
	ID        uint64
	CreatedAt int64  `gorm:"created_at"`
	UpdatedAt int64  `gorm:"updated_at"`
	DeletedAt int64  `gorm:"deleted_at"`
	UUID      string `gorm:"uuid"`
	Stage     string `gorm:"stage"`
}

func (inst *Instance) TableName() string {
	return "instance_tab"
}

func (inst *Instance) Create(db *gorm.DB) error {
	result := db.Create(inst)
	if result.Error != nil {
		return fmt.Errorf("create instance failed: %w", result.Error)
	}
	return nil
}
