package nomanssky

import "fmt"

type EntityCount[C NMSEntity] struct {
	Amount   int32  `yaml:"amount"`
	EntityID string `yaml:"id"`
}

func (ec EntityCount[C]) GetEntity(c []C) (*C, error) {
	for _, entity := range c {
		if ec.EntityID == entity.GetID() {
			entity := entity
			return &entity, nil
		}
	}
	return nil, fmt.Errorf("Building Not Found")
}
