package main

import (
	"fmt"
	"os"

	nomanssky "github.com/BeConfused/nms-planner-cli/internal/pkg/no-mans-sky"
	"github.com/BeConfused/nms-planner-cli/internal/pkg/plan"

	"gopkg.in/yaml.v3"
)

var configPaths []string

func main() {
	configPaths = []string{
		os.Getenv("CONFIG_FILE"),
		os.Getenv("PLAN_FILE"),
	}

	b := new(plan.Base)
	c := new(nomanssky.Config)

	// fmt.Println("Reading Config...")
	cf, mfErr := os.ReadFile(configPaths[0])
	if mfErr != nil {
		panic(mfErr)
	}

	// fmt.Println("Reading Plan...")
	pf, pfErr := os.ReadFile(configPaths[1])
	if pfErr != nil {
		panic(pfErr)
	}

	// fmt.Println("Loading Config...")
	umErr := yaml.Unmarshal(cf, c)
	if umErr != nil {
		panic(umErr)
	}

	fmt.Println("Loading Plan...")
	umErr = yaml.Unmarshal(pf, b)
	if umErr != nil {
		panic(umErr)
	}
	b.NMSConfig = *c

	report, rErr := b.GetReport(*c)
	if rErr != nil {
		panic(rErr)
	}
	fmt.Printf("%v", report.FormatAsString())
}
