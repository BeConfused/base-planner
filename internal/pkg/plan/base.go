package plan

import (
	"fmt"
	nomanssky "nms-planner-cli/internal/pkg/no-mans-sky"
)

type Base struct {
	Name      string                                      `yaml:"baseName"`
	Buildings []nomanssky.EntityCount[nomanssky.Building] `yaml:"buildings"`
}

func (b Base) Eval(c nomanssky.Config) (string, error) {
	report := fmt.Sprintf("Planned Base: %s\n", b.Name)
	report += fmt.Sprintf("Buildings: %s\n", "")

	for _, item := range b.Buildings {
		building, err := item.GetEntity(c.Buildings)
		if err != nil {
			return "", err
		}
		report += fmt.Sprintf("- %v * %s\n", item.Amount, building.Name)
	}

	return report, nil
}
