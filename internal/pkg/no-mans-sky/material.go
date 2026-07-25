package nomanssky

type Material struct {
	ID   string `yaml:"id"`
	Name string `yaml:"name"`
}

func (m Material) GetID() string {
	return m.ID
}

func (m Material) GetName() string {
	return m.Name
}
