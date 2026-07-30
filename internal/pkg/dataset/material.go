package dataset

// Material implements the Entity interface and provides a set of YAML annotations for unmarshaling
// when loaded from the given YAML Configuration. Unlike Building, it is also used as an intermediary.
type Material struct {
	ID   string `yaml:"id"`
	Name string `yaml:"name"`
}

// GetID is an implementation requirement from Entity: Returns the material's ID.
func (m Material) GetID() string {
	return m.ID
}

// GetName is an implementation requirement from Entity: Returns the material's Name.
func (m Material) GetName() string {
	return m.Name
}
