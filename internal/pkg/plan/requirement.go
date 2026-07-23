package plan

type Requirement[C any] struct {
	Requirement C     `yaml:",inline"`
	Amount      int32 `yaml:"amount"`
}
