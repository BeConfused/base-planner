package nomanssky

import "fmt"

type Config struct {
	Materials []Material `yaml:"materials"`
	Buildings []Building `yaml:"buildings"`
	Recipes   []Recipe   `yaml:"recipes"`
}

func (c Config) FindRecipe(eID string) (*Recipe, error) {
	for _, recipe := range c.Recipes {
		if eID == recipe.Output.EntityID {
			return &recipe, nil
		}
	}
	return nil, fmt.Errorf("No Recipe Found")
}
