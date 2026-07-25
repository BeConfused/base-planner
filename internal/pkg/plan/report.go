package plan

import (
	"fmt"
	"strings"

	nomanssky "github.com/BeConfused/nms-planner-cli/internal/pkg/no-mans-sky"
)

type Report[C nomanssky.NMSEntity] struct {
	List []Requirement[C]
}

func (r *Report[C]) FormatAsString() string {
	indentChar := ' '
	charsPerIndent := 2
	startingIndentLevel := 1

	var report strings.Builder

	report.WriteString("Report:\n")

	for _, item := range r.List {
		target := *item.Target
		report.WriteString(fmt.Sprintf("- %v x %s\n", item.Amount, target.GetName()))

		for _, material := range item.Materials {
			report.WriteString(indent(indentChar, charsPerIndent, startingIndentLevel))
			report.WriteString(fmt.Sprintf("- %v x %s\n", material.Amount, material.Target.GetName()))
		}
	}

	total := r.getTotal()

	report.WriteString("Total:\n")

	for _, item := range total {
		target := *item.Target
		report.WriteString(fmt.Sprintf("- %v x %s\n", item.Amount, target.GetName()))

		for _, material := range item.Materials {
			report.WriteString(indent(indentChar, charsPerIndent, startingIndentLevel))
			report.WriteString(fmt.Sprintf("- %v x %s\n", material.Amount, material.Target.GetName()))
		}
	}

	return report.String()
}

func (r *Report[C]) getTotal() []Requirement[nomanssky.Material] {
	total := []Requirement[nomanssky.Material]{}

	for _, item := range r.List {
		for _, material := range item.Materials {
			foundOn := findIndexByID(total, material)
			if foundOn == -1 {
				total = append(total, material)
			} else {
				total[foundOn].Amount += material.Amount
			}
		}
	}

	return total
}

func indent(s rune, rpl int, l int) string {
	return strings.Repeat(string(s), l*rpl)
}

func findIndexByID[C nomanssky.NMSEntity](array []Requirement[C], f Requirement[C]) int {
	toFind := *f.Target
	for index, item := range array {
		current := *item.Target
		if toFind.GetID() == current.GetID() {
			return index
		}
	}

	return -1
}
