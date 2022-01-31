package instance

import (
	"context"
	"fmt"

	"github.com/berryfl/internal/helper"
	"github.com/berryfl/event/model/po"
	"github.com/berryfl/event/model/vo"
)

const (
	NewInstanceList = "new_instance_list"
)

func CreateInstance(inst *vo.Instance) error {
	db := helper.GetDB()
	rdb := helper.GetRedis()

	poInst := inst.ToPoInstance()
	if err := poInst.Create(db); err != nil {
		return fmt.Errorf("create instance failed: %w", err)
	}

	data, err := inst.Serialize()
	if err != nil {
		return fmt.Errorf("create instance failed: %w", err)
	}

	cmd := rdb.LPush(context.Background(), NewInstanceList, data)
	if _, err := cmd.Result(); err != nil {
		return fmt.Errorf("create instance failed: %w", err)
	}

	return nil
}
