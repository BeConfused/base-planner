package nomanssky

import (
	"errors"
	"fmt"
)

type Config struct {
	Materials []Material `yaml:"materials"`
	Buildings []Building `yaml:"buildings"`
	Recipes   []Recipe   `yaml:"recipes"`
}

var ErrRecipeNotFound = errors.New("no recipe found")

func (c Config) FindRecipe(eID string) (*Recipe, error) {
	for _, recipe := range c.Recipes {
		if eID == recipe.Output.EntityID {
			return &recipe, nil
		}
	}
	return nil, fmt.Errorf("%w: Ensure, that the recipes Esxist and all IDs are properly set", ErrRecipeNotFound)
}
