package nomanssky

type Config struct {
	Materials []Material `yaml:"materials"`
	Buildings []Building `yaml:"buildings"`
	Recipes   []Recipe   `yaml:"recipes"`
}
