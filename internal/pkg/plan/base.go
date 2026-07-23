package plan

import nomanssky "nms-planner-cli/internal/pkg/no-mans-sky"

type Base struct {
	Name      string                            `yaml:"baseName"`
	Buildings []Requirement[nomanssky.Building] `yaml:"buildings"`
}
