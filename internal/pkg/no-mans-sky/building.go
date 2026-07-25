package nomanssky

// Building implements the NMSEntity interface and provides a set of YAML annotations for unmarshaling
// when loaded from the given YAML Configuration.
type Building struct {
	ID   string `yaml:"id"`
	Name string `yaml:"name"`
}

// GetID is an implementation requirement from NMS Entity: Returns the building's ID.
func (b Building) GetID() string {
	return b.ID
}

// GetName is an implementation requirement from NMSEntity: Returns the building's name.
func (b Building) GetName() string {
	return b.Name
}
