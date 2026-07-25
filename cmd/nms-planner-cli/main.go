package main

import (
	"fmt"
	"os"
	"path/filepath"

	nomanssky "github.com/BeConfused/nms-planner-cli/internal/pkg/no-mans-sky"
	"github.com/BeConfused/nms-planner-cli/internal/pkg/plan"
	"github.com/BeConfused/nms-planner-cli/internal/pkg/util"

	"gopkg.in/yaml.v3"
)

func main() {
	paths := util.PathList{
		Config: os.Getenv("CONFIG_FILE"),
		Plan:   os.Getenv("PLAN_FILE"),
	}

	base := new(plan.Base)
	config := new(nomanssky.Config)

	// fmt.Println("Reading Config...")
	cleanConfigPath := filepath.Clean(paths.Config)
	configFile, configFileErr := os.ReadFile(cleanConfigPath)
	if configFileErr != nil {
		panic(configFileErr)
	}

	// fmt.Println("Reading Plan...")
	cleanPlanPath := filepath.Clean(paths.Plan)
	pathFile, pathFileErr := os.ReadFile(cleanPlanPath)
	if pathFileErr != nil {
		panic(pathFileErr)
	}

	// fmt.Println("Loading Config...")
	umErr := yaml.Unmarshal(configFile, config)
	if umErr != nil {
		panic(umErr)
	}

	fmt.Println("Loading Plan...")
	umErr = yaml.Unmarshal(pathFile, base)
	if umErr != nil {
		panic(umErr)
	}
	base.NMSConfig = *config

	report, rErr := base.GetReport(*config)
	if rErr != nil {
		panic(rErr)
	}
	fmt.Printf("%v", report.FormatAsString())
}
