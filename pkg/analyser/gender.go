package analyser

import (
	"slices"

	"github.com/basbeu/JudobaseStats/internal/category"
	"github.com/basbeu/JudobaseStats/internal/judobase"
)

type gender int

const (
	male gender = iota
	female
	unknownGender
)

var genders []gender

func init() {
	genders = []gender{male, female}
}

func (g gender) string() string {
	switch g {
	case male:
		return "Male"
	case female:
		return "Female"
	default:
		return "Unknown"
	}
}

func parseGender(contest judobase.Contest) gender {
	if contest.IDWeight != nil {
		cat, err := category.FromWeightID(*contest.IDWeight)
		if err != nil {
			return unknownGender
		}
		if slices.Contains(category.MaleCategories, cat) {
			return male
		}
		if slices.Contains(category.FemaleCategories, cat) {
			return female
		}
	}
	return unknownGender
}
