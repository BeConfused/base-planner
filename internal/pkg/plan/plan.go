package plan

import nomanssky "nms-planner-cli/internal/pkg/no-mans-sky"

type Plan interface {
	Eval(nomanssky.Config) error
}
