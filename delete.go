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
	if err := po.DeleteCreated(db); err != nil {
		panic(err)
	}
}
