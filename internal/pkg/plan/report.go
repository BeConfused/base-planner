package plan

import (
	"fmt"
	"strings"

	"github.com/BeConfused/base-planner/internal/pkg/dataset"
)

// Report eceives a List of Requirements to Report on and exposes
// a set of Functions to receive them back in a formatted way.
type Report[C dataset.Entity] struct {
	List []Requirement[C]
}

// FormatAsString returns the report as a concatenated string.
func (r *Report[C]) FormatAsString() string {
	indentChar := ' '
	charsPerIndent := 2
	startingIndentLevel := 1

	var report strings.Builder

	fmt.Fprint(&report, "Report:\n")

	for _, item := range r.List {
		target := *item.Target
		fmt.Fprintf(&report, "- %v x %s\n", item.Amount, target.GetName())

		for _, material := range item.Materials {
			fmt.Fprint(&report, indent(indentChar, charsPerIndent, startingIndentLevel))
			fmt.Fprintf(&report, "- %v x %s\n", material.Amount, material.Target.GetName())
		}
	}

	total := r.getTotal()

	fmt.Fprint(&report, "Total:\n")

	for _, item := range total {
		target := *item.Target
		fmt.Fprintf(&report, "- %v x %s\n", item.Amount, target.GetName())

		for _, material := range item.Materials {
			fmt.Fprint(&report, indent(indentChar, charsPerIndent, startingIndentLevel))
			fmt.Fprintf(&report, "- %v x %s\n", material.Amount, material.Target.GetName())
		}
	}

	return report.String()
}

/* Summarizes all materials required by its list of entities.
 * For Example:
 * material a is needed by entity 1 and entity 3.
 * it will add the material to the list on entity 1 and then increment its amount by the amount of entity 3.
 */
func (r *Report[C]) getTotal() []Requirement[dataset.Material] {
	total := []Requirement[dataset.Material]{}

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

func findIndexByID[C dataset.Entity](array []Requirement[C], f Requirement[C]) int {
	toFind := *f.Target
	for index, item := range array {
		current := *item.Target
		if toFind.GetID() == current.GetID() {
			return index
		}
	}

	return -1
}
