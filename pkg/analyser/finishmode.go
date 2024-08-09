package analyser

import "github.com/basbeu/JudobaseStats/internal/judobase"

type finishMode bool

const (
	goldenScore finishMode = true
	regularTime finishMode = false
)

func (g finishMode) string() string {
	if g {
		return "Golden score"
	}
	return "Regular time"
}

func parseFinishMode(contest judobase.Contest) finishMode {
	return contest.GoldenScore != nil && *contest.GoldenScore == "1"
}
