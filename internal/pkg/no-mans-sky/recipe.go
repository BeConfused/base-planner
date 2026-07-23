package nomanssky

type Recipe struct {
	Output RecipeComponent   `yaml:",inline"`
	Input  []RecipeComponent `yaml:"requirements"`
}
