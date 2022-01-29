package main

import (
	"github.com/berryfl/event/internal/helper"
	"github.com/berryfl/event/model/po"
)

func main() {
	if err := helper.InitDB(); err != nil {
		panic(err)
	}

	db := helper.GetDB()
	instances, err := po.QueryNew(db)
	if err != nil {
		panic(err)
	}
	for _, inst := range instances {
		inst.Stage = "CREATED"
		if err := inst.Save(db); err != nil {
			panic(err)
		}
	}
}
