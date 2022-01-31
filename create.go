package main

import (
	"github.com/berryfl/event/internal/instance"
	"github.com/berryfl/event/internal/helper"
	"github.com/berryfl/event/model/vo"
	"github.com/google/uuid"
)

func main() {
	if err := helper.InitDB(); err != nil {
		panic(err)
	}
	helper.InitRedis()

	inst1 := &vo.Instance{
		UUID:  uuid.New().String(),
		Stage: "NEW",
	}
	inst2 := &vo.Instance{
		UUID:  uuid.New().String(),
		Stage: "NEW",
	}
	inst3 := &vo.Instance{
		UUID:  uuid.New().String(),
		Stage: "NEW",
	}

	if err := instance.CreateInstance(inst1); err != nil {
		panic(err)
	}
	if err := instance.CreateInstance(inst2); err != nil {
		panic(err)
	}
	if err := instance.CreateInstance(inst3); err != nil {
		panic(err)
	}
}
