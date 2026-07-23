package plan

type Plan interface {
	Eval() error
}
