package instance

import (
	"fmt"

	"github.com/berryfl/event/model/po"
	"github.com/berryfl/event/model/vo"
)

func HandleNewInstance(inst *vo.Instance) error {
	fmt.Printf("simulating work to handle new instance: %s\n", inst.UUID)

	poInst := inst.ToPoInstance()
	if err := poInst.UpdateStage(); err != nil {
		return fmt.Errorf("handle new instance failed: %w", err)
	}
	return nil
}
