package dataset

// Entity is an interface to guarantee a set of available functions.
type Entity interface {
	GetID() string
	GetName() string
}
