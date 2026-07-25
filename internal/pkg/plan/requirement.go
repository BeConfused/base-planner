package plan

import (
	"fmt"

	nomanssky "github.com/BeConfused/nms-planner-cli/internal/pkg/no-mans-sky"
)

// Requirement is a generic type, that is supposed to function as a resolved
// EntityCount with direct access to a given entity.
type Requirement[C nomanssky.NMSEntity] struct {
	Target    *C
	Amount    int32
	Materials []Requirement[nomanssky.Material]
}

// FromEntityCount creates a Requirement given an NMSEntity.
// Invokes an internal recusive implementation with a starting value of 1.
// Parameters:
// - config: Given to find a recipe to resolve against.
// - c: a list of entities to resolve the recipe against.
// - ec: the entityCount to create the Requirement from.
func FromEntityCount[C nomanssky.NMSEntity](
	config nomanssky.Config,
	c []C,
	ec nomanssky.EntityCount[C],
) (*Requirement[C], error) {
	return fromEntityCount(config, c, ec, 1)
}

func fromEntityCount[C nomanssky.NMSEntity](
	config nomanssky.Config,
	c []C,
	entityCount nomanssky.EntityCount[C],
	amountAmp int32,
) (*Requirement[C], error) {
	entityRef, geErr := entityCount.GetEntity(c)
	if geErr != nil {
		return nil, fmt.Errorf("building requirement failed: %w", geErr)
	}

	entity := *entityRef

	materials := []Requirement[nomanssky.Material]{}

	recipe, rErr := config.FindRecipe(entity.GetID())
	if rErr == nil {
		for _, input := range recipe.Input {
			material, mErr := fromEntityCount(config, config.Materials, input, entityCount.Amount)
			if mErr != nil {
				return nil, mErr
			}

			materials = append(materials, *material)
		}
	}

	return &Requirement[C]{
		Target:    entityRef,
		Amount:    entityCount.Amount * amountAmp,
		Materials: materials,
	}, nil
}
