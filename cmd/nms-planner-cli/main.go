package main

import (
	"fmt"
	"nms-planner-cli/internal/pkg/plan"
	"nms-planner-cli/internal/pkg/util"
	"os"

	"gopkg.in/yaml.v3"
)

var configPaths []string

func init() {
	configPaths = []string{
		os.Getenv("CONFIG_FILE"),
		os.Getenv("PLAN_FILE"),
	}
}

func main() {
	b := &plan.Base{}
	c := &util.Config{}

	fmt.Println("Reading Config...")
	cf, mfErr := os.ReadFile(configPaths[0])
	if mfErr != nil {
		panic(mfErr)
	}

	fmt.Println("Reading Plan...")
	pf, pfErr := os.ReadFile(configPaths[1])
	if pfErr != nil {
		panic(pfErr)
	}

	fmt.Println("Loading Config...")
	yaml.Unmarshal(cf, c)
	fmt.Printf("%v\n", c)

	fmt.Println("Loading Plan...")
	yaml.Unmarshal(pf, b)
	fmt.Printf("%v\n", b)
}
