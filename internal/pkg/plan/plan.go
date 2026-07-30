// Package plan provides the basic functionality to manage a given plan and creating a summary
package plan

import dataset "github.com/BeConfused/base-planner/internal/pkg/dataset"

// Plan exposes a set of functions to guarantee a report, based on a given plan.
type Plan interface {
	GetReport(c dataset.Config) Report[dataset.Entity]
}
