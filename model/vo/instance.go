package vo

import (
	"encoding/json"
	"fmt"

	"github.com/berryfl/event/model/po"
)

type Instance struct {
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	UUID      string `json:"uuid"`
	Stage     string `json:"stage"`
}

func FromPoInstance(poInst *po.Instance) *Instance {
	return &Instance{
		CreatedAt: poInst.CreatedAt,
		UpdatedAt: poInst.UpdatedAt,
		UUID:      poInst.UUID,
		Stage:     poInst.Stage,
	}
}

func (inst *Instance) ToPoInstance() *po.Instance {
	return &po.Instance{
		UUID:  inst.UUID,
		Stage: inst.Stage,
	}
}

func FromJSONInstance(data []byte) (*Instance, error) {
	var inst Instance
	if err := json.Unmarshal(data, &inst); err != nil {
		return nil, fmt.Errorf("deserialize instance failed: %w", err)
	}
	return &inst, nil
}

func (inst *Instance) Serialize() ([]byte, error) {
	result, err := json.Marshal(inst)
	if err != nil {
		return nil, fmt.Errorf("serialize instance failed: %w", err)
	}
	return result, nil
}
