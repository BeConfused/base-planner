package main

import (
	"fmt"
	"os"

	nomanssky "github.com/BeConfused/nms-planner-cli/internal/pkg/no-mans-sky"
	"github.com/BeConfused/nms-planner-cli/internal/pkg/plan"

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
	yaml.Unmarshal(cf, c)

	fmt.Println("Loading Plan...")
	yaml.Unmarshal(pf, b)
	b.NMSConfig = *c

	report, rErr := b.GetReport(*c)
	if rErr != nil {
		panic(rErr)
	}
	fmt.Printf("%v", report.FormatAsString())
}
