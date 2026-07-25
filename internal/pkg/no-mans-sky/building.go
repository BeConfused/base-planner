package nomanssky

type Building struct {
	ID   string `yaml:"id"`
	Name string `yaml:"name"`
}

func (b Building) GetID() string {
	return b.ID
}

func (b Building) GetName() string {
	return b.Name
}
