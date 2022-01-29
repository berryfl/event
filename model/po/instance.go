package po

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type Instance struct {
	ID        uint64
	CreatedAt int64                 `gorm:"created_at"`
	UpdatedAt int64                 `gorm:"updated_at"`
	DeletedAt soft_delete.DeletedAt `gorm:"deleted_at"`
	UUID      string                `gorm:"uuid"`
	Stage     string                `gorm:"stage"`
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

func (inst *Instance) Save(db *gorm.DB) error {
	result := db.Save(inst)
	if result.Error != nil {
		return fmt.Errorf("save instance failed: %w", result.Error)
	}
	return nil
}

func QueryNew(db *gorm.DB) ([]*Instance, error) {
	var instances []*Instance
	result := db.Where(&Instance{Stage: "NEW"}).Find(&instances)
	if result.Error != nil {
		return nil, fmt.Errorf("query stage new instances failed: %w", result.Error)
	}
	return instances, nil
}

func DeleteCreated(db *gorm.DB) error {
	result := db.Where(&Instance{Stage: "CREATED"}).Delete(&Instance{})
	if result.Error != nil {
		return fmt.Errorf("delete created instances failed: %w", result.Error)
	}
	return nil
}
