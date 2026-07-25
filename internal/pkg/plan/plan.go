// Package plan provides the basic functionality to manage a given plan and creating a summary
package plan

import nomanssky "github.com/BeConfused/nms-planner-cli/internal/pkg/no-mans-sky"

// Plan exposes a set of functions to guarantee a report, based on a given plan.
type Plan interface {
	GetReport(c nomanssky.Config) Report[nomanssky.NMSEntity]
}
