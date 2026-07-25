package plan

import (
	nomanssky "github.com/BeConfused/nms-planner-cli/internal/pkg/no-mans-sky"
)

type Base struct {
	NMSConfig nomanssky.Config                            `yaml:"customConfig,omitempty"`
	Name      string                                      `yaml:"baseName"`
	Buildings []nomanssky.EntityCount[nomanssky.Building] `yaml:"buildings"`
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
