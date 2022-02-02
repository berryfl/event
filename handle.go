package main

import (
	"github.com/berryfl/event/internal/instance"
	"github.com/berryfl/event/internal/helper"
)

func main() {
	if err := helper.InitDB(); err != nil {
		panic(err)
	}
	helper.InitRedis()

	if err := instance.LoopNewInstance(); err != nil {
		panic(err)
	}
}
