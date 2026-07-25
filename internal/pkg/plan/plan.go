package plan

import nomanssky "github.com/BeConfused/nms-planner-cli/internal/pkg/no-mans-sky"

type Plan interface {
	GetReport(nomanssky.Config) Report[nomanssky.Building]
}
