package nomanssky

// NMSEntity is an interface to guarantee a set of available functions.
type NMSEntity interface {
	GetID() string
	GetName() string
}
