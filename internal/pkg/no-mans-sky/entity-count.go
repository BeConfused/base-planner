package nomanssky

import (
	"errors"
	"fmt"
)

var ErrEntityNotFound = errors.New("Entity not found")

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
	return nil, fmt.Errorf("%w: Ensure, that an entity of type %T exists", ErrEntityNotFound, new(C))
}
