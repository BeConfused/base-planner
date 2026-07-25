package nomanssky

type Recipe struct {
	Output EntityCount[NMSEntity]  `yaml:",inline"`
	Input  []EntityCount[Material] `yaml:"requirements"`
}
