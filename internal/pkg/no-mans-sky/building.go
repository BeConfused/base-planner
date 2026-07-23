package nomanssky

type Building struct {
	ID   string `yaml:"id"`
	Name string `yaml:"name"`
}

func (b Building) GetID() string {
	return b.ID
}
