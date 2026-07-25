// Package nomanssky provides all functions required to interact with
// a given list of entities, based on a loaded yaml configuration.
package nomanssky

import (
	"errors"
	"fmt"
)

// Config is the base type to load the original YAML-Configuration from.
type Config struct {
	Materials []Material `yaml:"materials"`
	Buildings []Building `yaml:"buildings"`
	Recipes   []Recipe   `yaml:"recipes"`
}

// ErrRecipeNotFound is a static error, that will be returned in case no recipe could be found by FindRecipe.
var ErrRecipeNotFound = errors.New("no recipe found")

// FindRecipe looks for a Recipe given a specific EntityID.
func (c Config) FindRecipe(eID string) (*Recipe, error) {
	for _, recipe := range c.Recipes {
		if eID == recipe.Output.EntityID {
			return &recipe, nil
		}
	}

	return nil, fmt.Errorf("%w: Ensure, that the recipes exist and all IDs are properly set", ErrRecipeNotFound)
}

// IsEmpty checks if the configuration is completely empty.
// Implies an invalid configuration.
func (c Config) IsEmpty() bool {
	return len(c.Materials) == 0 && len(c.Buildings) == 0 && len(c.Recipes) == 0
}
