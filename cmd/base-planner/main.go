// package main provides the basic entrypoint for the CLI
package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BeConfused/base-planner/internal/pkg/dataset"
	"github.com/BeConfused/base-planner/internal/pkg/plan"
	"github.com/BeConfused/base-planner/internal/pkg/util"

	"gopkg.in/yaml.v3"
)

func main() {
	paths := util.PathList{
		Config: os.Getenv("CONFIG_FILE"),
		Plan:   os.Getenv("PLAN_FILE"),
	}

	base := new(plan.Base)
	config := new(dataset.Config)

	cleanConfigPath := filepath.Clean(paths.Config)

	configFile, configFileErr := os.ReadFile(cleanConfigPath)
	if configFileErr != nil {
		panic(configFileErr)
	}

	cleanPlanPath := filepath.Clean(paths.Plan)

	pathFile, pathFileErr := os.ReadFile(cleanPlanPath)
	if pathFileErr != nil {
		panic(pathFileErr)
	}

	umErr := yaml.Unmarshal(configFile, config)
	if umErr != nil {
		panic(umErr)
	}

	umErr = yaml.Unmarshal(pathFile, base)
	if umErr != nil {
		panic(umErr)
	}

	if base.DataConfig.IsEmpty() {
		base.DataConfig = *config
	} else { // Temporary fix: Larger refactoring required
		config = &base.DataConfig
	}

	report, rErr := base.GetReport(*config)
	if rErr != nil {
		panic(rErr)
	}

	_, err := fmt.Fprintf(os.Stdout, "%v", report.FormatAsString())
	if err != nil {
		panic(err)
	}
}
