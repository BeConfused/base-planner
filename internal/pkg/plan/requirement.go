package plan

import nomanssky "nms-planner-cli/internal/pkg/no-mans-sky"

type Requirement[C nomanssky.NMSEntity] struct {
	Target    *C
	Amount    int32
	Materials []Requirement[nomanssky.Material]
}

func FromEntityCount[C nomanssky.NMSEntity](config nomanssky.Config, c []C, ec nomanssky.EntityCount[C], amountAmp int32) (*Requirement[C], error) {
	entityRef, geErr := ec.GetEntity(c)
	if geErr != nil {
		return nil, geErr
	}
	entity := *entityRef

	materials := []Requirement[nomanssky.Material]{}
	recipe, rErr := config.FindRecipe(entity.GetID())
	if rErr == nil {
		for _, input := range recipe.Input {
			material, mErr := FromEntityCount(config, config.Materials, input, ec.Amount)
			if mErr != nil {
				return nil, mErr
			}
			materials = append(materials, *material)
		}
	}

	return &Requirement[C]{
		Target:    entityRef,
		Amount:    ec.Amount * amountAmp,
		Materials: materials,
	}, nil
}
