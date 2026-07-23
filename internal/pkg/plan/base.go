package plan

import (
	"fmt"
	nomanssky "nms-planner-cli/internal/pkg/no-mans-sky"
)

type Base struct {
	Name      string                            `yaml:"baseName"`
	Buildings []Requirement[nomanssky.Building] `yaml:"buildings"`
}

func (b Base) Eval() error {
	return fmt.Errorf("%s not yet Implemented", "Base.Eval()")
}
