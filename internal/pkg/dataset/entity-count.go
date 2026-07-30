package dataset

import (
	"errors"
	"fmt"
)

// ErrEntityNotFound is a static error, that will be returned in case no entity could be found by GetEntity.
var ErrEntityNotFound = errors.New("entity not found")

// EntityCount loads an unresolved id of a given entity, that can be Resolved using the GetEntity method.
type EntityCount[C Entity] struct {
	Amount   int32  `yaml:"amount"`
	EntityID string `yaml:"id"`
}

// GetEntity looks up a given Entity based on the EntityCount's given ID. Requires an array to check against
// for now.
func (ec EntityCount[C]) GetEntity(c []C) (*C, error) {
	for _, entity := range c {
		if ec.EntityID == entity.GetID() {
			entity := entity

			return &entity, nil
		}
	}

	return nil, fmt.Errorf("%w: Ensure, that an entity of type %T exists", ErrEntityNotFound, new(C))
}
