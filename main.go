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
		os.Getenv("PLAN_FILE"),
	}
}

func main() {
	fmt.Printf("hello, %s\n", "world!")
	b := &plan.Base{}
	file, err := os.ReadFile(configPaths[0])
	if err != nil {
		panic(err)
	}
	yaml.Unmarshal(file, b)
	fmt.Printf("%s\n", file)
	fmt.Printf("%v\n", b)
}
