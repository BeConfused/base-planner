package plan

import (
	nomanssky "BeConfused/nms-planner-cli/internal/pkg/no-mans-sky"
	"fmt"
)

type Base struct {
	NMSConfig nomanssky.Config
	Name      string                                      `yaml:"baseName"`
	Buildings []nomanssky.EntityCount[nomanssky.Building] `yaml:"buildings"`
}

func (b Base) Eval() (string, error) {
	report := fmt.Sprintf("Planned Base: %s\n", b.Name)
	report += fmt.Sprintf("Buildings: %s\n", "")

	for _, item := range b.Buildings {
		building, err := item.GetEntity(b.NMSConfig.Buildings)
		if err != nil {
			return "", err
		}
		report += fmt.Sprintf("- %v * %s\n", item.Amount, building.Name)
		recipe, err := b.NMSConfig.FindRecipe(item.EntityID)
		if err != nil {
			return "", err
		}

		for _, material := range recipe.Input {
			report += fmt.Sprintf("  - %v * %s\n", material.Amount*item.Amount, material.EntityID)
		}
	}

	return report, nil
}

func (b Base) GetReport(c nomanssky.Config) (*Report[nomanssky.Building], error) {
	buildings := []Requirement[nomanssky.Building]{}
	for _, item := range b.Buildings {
		building, bErr := FromEntityCount(c, b.NMSConfig.Buildings, item, 1)
		if bErr != nil {
			return nil, bErr
		}
		buildings = append(buildings, *building)
	}
	return &Report[nomanssky.Building]{
		List: buildings,
	}, nil
}
