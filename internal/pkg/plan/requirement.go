package plan

import (
	"fmt"

	"github.com/BeConfused/base-planner/internal/pkg/dataset"
)

// Requirement is a generic type, that is supposed to function as a resolved
// EntityCount with direct access to a given entity.
type Requirement[C dataset.Entity] struct {
	Target    *C
	Amount    int32
	Materials []Requirement[dataset.Material]
}

// FromEntityCount creates a Requirement given an Entity.
// Invokes an internal recursive implementation with a starting value of 1.
// Parameters:
// - config: Given to find a recipe to resolve against.
// - c: a list of entities to resolve the recipe against.
// - ec: the entityCount to create the Requirement from.
func FromEntityCount[C dataset.Entity](
	config dataset.Config,
	c []C,
	ec dataset.EntityCount[C],
) (*Requirement[C], error) {
	return fromEntityCount(config, c, ec, 1)
}

func fromEntityCount[C dataset.Entity](
	config dataset.Config,
	c []C,
	entityCount dataset.EntityCount[C],
	amountAmp int32,
) (*Requirement[C], error) {
	entityRef, geErr := entityCount.GetEntity(c)
	if geErr != nil {
		return nil, fmt.Errorf("building requirement failed: %w", geErr)
	}

	entity := *entityRef

	materials := []Requirement[dataset.Material]{}

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
