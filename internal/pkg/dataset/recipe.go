package dataset

// Recipe defines a list of Required materials for a specific entity.
// Caution: This has only been manually tested to work with Building as an output!
type Recipe struct {
	Output EntityCount[Entity]     `yaml:",inline"`
	Input  []EntityCount[Material] `yaml:"requirements"`
}
