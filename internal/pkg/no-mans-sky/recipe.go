package nomanssky

type Recipe struct {
	Output EntityCount[NMSEntity]   `yaml:",inline"`
	Input  []EntityCount[NMSEntity] `yaml:"requirements"`
}
