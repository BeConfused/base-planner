package plan

import (
	"github.com/BeConfused/base-planner/internal/pkg/dataset"
)

// Base implements the plan interface to create a report based on an inventory of building references.
type Base struct {
	DataConfig dataset.Config                          `yaml:"customConfig,omitempty"`
	Name       string                                  `yaml:"baseName"`
	Buildings  []dataset.EntityCount[dataset.Building] `yaml:"buildings"`
}

// GetReport is an implementation of the plan interface. it resolves all given entity counts and adds them
// to a list on a report for output.
func (b Base) GetReport(c dataset.Config) (*Report[dataset.Building], error) {
	buildings := []Requirement[dataset.Building]{}

	for _, item := range b.Buildings {
		building, bErr := FromEntityCount(c, b.DataConfig.Buildings, item)
		if bErr != nil {
			return nil, bErr
		}

		buildings = append(buildings, *building)
	}

	return &Report[dataset.Building]{
		List: buildings,
	}, nil
}
