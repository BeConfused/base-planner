package main

import (
	"fmt"
	"os"

	nomanssky "github.com/BeConfused/nms-planner-cli/internal/pkg/no-mans-sky"
	"github.com/BeConfused/nms-planner-cli/internal/pkg/plan"

	"gopkg.in/yaml.v3"
)

func main() {
	configPaths := []string{
		os.Getenv("CONFIG_FILE"),
		os.Getenv("PLAN_FILE"),
	}

	base := new(plan.Base)
	config := new(nomanssky.Config)

	// fmt.Println("Reading Config...")
	configFile, configFileErr := os.ReadFile(configPaths[0])
	if configFileErr != nil {
		panic(configFileErr)
	}

	// fmt.Println("Reading Plan...")
	pathFile, pathFileErr := os.ReadFile(configPaths[1])
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
