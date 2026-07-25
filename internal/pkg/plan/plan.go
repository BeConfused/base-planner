package plan

import nomanssky "nms-planner-cli/internal/pkg/no-mans-sky"

type Plan interface {
	GetReport(nomanssky.Config) Report[nomanssky.Building]
}
