package main

import (
	"fmt"
	"nms-planner-cli/internal/pkg/plan"
	"os"

	"gopkg.in/yaml.v3"
)

var configPaths []string

func init() {
	configPaths = []string{
		os.Getenv("MATERIALS_FILE"),
		os.Getenv("RECIPE_FILE"),
		os.Getenv("PLAN_FILE"),
	}
}

func main() {
	b := &plan.Base{}

	_, mfErr := os.ReadFile(configPaths[0])
	if mfErr != nil {
		panic(mfErr)
	}

	_, rfErr := os.ReadFile(configPaths[1])
	if rfErr != nil {
		panic(rfErr)
	}

	pf, pfErr := os.ReadFile(configPaths[2])
	if pfErr != nil {
		panic(pfErr)
	}

	yaml.Unmarshal(pf, b)
	fmt.Printf("%s\n", pf)
	fmt.Printf("%v\n", b)
}
