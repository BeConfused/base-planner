package main

import (
	"fmt"
	nomanssky "nms-planner-cli/internal/pkg/no-mans-sky"
	"nms-planner-cli/internal/pkg/plan"
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
	c := &nomanssky.Config{}

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
	report, err := b.Eval(*c)
	if err != nil {
		panic(err)
	}
	fmt.Println(report)
}
