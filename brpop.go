package main

import (
	"context"
	"fmt"
	"time"

	"github.com/berryfl/event/internal/helper"
)

var (
	duration time.Duration
)

func main() {
	helper.InitRedis()
	rdb := helper.GetRedis()

	for {
		cmd := rdb.BRPop(context.Background(), duration, "eventlist")
		data, err := cmd.Result()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", data)
	}
}
