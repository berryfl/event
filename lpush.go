package main

import (
	"context"

	"github.com/berryfl/event/internal/helper"
	"github.com/google/uuid"
)

func main() {
	helper.InitRedis()
	rdb := helper.GetRedis()

	for i := 0; i < 3; i++ {
		value := uuid.New().String()
		cmd := rdb.LPush(context.Background(), "eventlist", value)
		if _, err := cmd.Result(); err != nil {
			panic(err)
		}
	}
}
