package main

import (
	"github.com/berryfl/event/internal/helper"
	"github.com/berryfl/event/model/po"
	"github.com/google/uuid"
)

func main() {
	if err := helper.InitDB(); err != nil {
		panic(err)
	}

	inst1 := &po.Instance{
		UUID:  uuid.New().String(),
		Stage: "NEW",
	}
	inst2 := &po.Instance{
		UUID:  uuid.New().String(),
		Stage: "NEW",
	}
	inst3 := &po.Instance{
		UUID:  uuid.New().String(),
		Stage: "NEW",
	}

	db := helper.GetDB()
	if err := inst1.Create(db); err != nil {
		panic(err)
	}
	if err := inst2.Create(db); err != nil {
		panic(err)
	}
	if err := inst3.Create(db); err != nil {
		panic(err)
	}
}
