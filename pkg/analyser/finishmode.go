package analyser

import "github.com/basbeu/JudobaseStats/internal/judobase"

type FinishMode bool

const (
	goldenScore FinishMode = true
	regularTime FinishMode = false
)

func (g FinishMode) String() string {
	if g {
		return "Golden score"
	}
	return "Regular time"
}

func parseFinishMode(contest judobase.Contest) FinishMode {
	return contest.GoldenScore != nil && *contest.GoldenScore == "1"
}
