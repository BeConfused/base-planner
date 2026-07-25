package plan

import (
	nomanssky "github.com/BeConfused/nms-planner-cli/internal/pkg/no-mans-sky"
)

// Base implements the plan interface to create a report based on an inventory of building references.
type Base struct {
	NMSConfig nomanssky.Config                            `yaml:"customConfig,omitempty"`
	Name      string                                      `yaml:"baseName"`
	Buildings []nomanssky.EntityCount[nomanssky.Building] `yaml:"buildings"`
}

// GetReport is an implementation of the plan interface. it resolves all given entity counts and adds them
// to a list on a report for output.
func (b Base) GetReport(c nomanssky.Config) (*Report[nomanssky.Building], error) {
	buildings := []Requirement[nomanssky.Building]{}

	for _, item := range b.Buildings {
		building, bErr := FromEntityCount(c, b.NMSConfig.Buildings, item)
		if bErr != nil {
			return nil, bErr
		}

		buildings = append(buildings, *building)
	}

	return &Report[nomanssky.Building]{
		List: buildings,
	}, nil
}
