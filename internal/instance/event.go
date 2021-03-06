package instance

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/berryfl/event/internal/helper"
	"github.com/berryfl/event/model/vo"
	"github.com/bsm/redislock"
)

const (
	duration time.Duration = 0
)

func LoopNewInstance() error {
	rdb := helper.GetRedis()
	for {
		cmd := rdb.BRPop(context.Background(), duration, NewInstanceList)
		data, err := cmd.Result()
		if err != nil {
			return fmt.Errorf("loop new instances failed: %w", err)
		}
		if len(data) < 2 {
			return fmt.Errorf("BRPop command should return key and value")
		}
		switch data[0] {
		case NewInstanceList:
			if err := HandleNewInstanceList(data[1:]); err != nil {
				return fmt.Errorf("loop new instances failed: %w", err)
			}
		}
	}
}

func HandleNewInstanceList(values []string) error {
	for _, value := range values {
		inst, err := vo.FromJSONInstance([]byte(value))
		if err != nil {
			return fmt.Errorf("handle new instance failed: %w", err)
		}
		if err := HandleNewInstance(inst); err != nil {
			return fmt.Errorf("handle new instance failed: %w", err)
		}
	}
	return nil
}

func HandleNewInstance(inst *vo.Instance) error {
	db := helper.GetDB()

	rdb := helper.GetRedis()
	locker := redislock.New(rdb)
	key := helper.GetLockKey(inst.UUID)
	ttl := 15 * time.Second
	lock, err := locker.Obtain(context.Background(), key, ttl, nil)
	if err != nil {
		if errors.Is(err, redislock.ErrNotObtained) {
			fmt.Printf("cannot obtain lock for new instance: %s\n", inst.UUID)
			return nil
		}
		return fmt.Errorf("handle new instance failed: %w", err)
	}
	defer lock.Release(context.Background())

	fmt.Printf("simulating work to handle new instance: %s\n", inst.UUID)
	inst.Stage = "CREATED"

	poInst := inst.ToPoInstance()
	if err := poInst.UpdateStage(db); err != nil {
		return fmt.Errorf("handle new instance failed: %w", err)
	}
	return nil
}
