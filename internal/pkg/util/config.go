package util

import nomanssky "nms-planner-cli/internal/pkg/no-mans-sky"

type Config struct {
	Materials []nomanssky.Material `yaml:"materials"`
	Buildings []nomanssky.Building `yaml:"buildings"`
	Recipes   *[]nomanssky.Recipe  `yaml:"recipes"`
}
