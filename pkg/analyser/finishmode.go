package analyser

import (
	"strconv"
	"strings"
	"time"

	"github.com/basbeu/JudobaseStats/internal/judobase"
)

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
	if contest.Duration != nil && contest.CompYear != nil {
		d, err := time.ParseDuration(strings.Replace(strings.Replace(*contest.Duration, ":", "h", 1), ":", "m", 1) + "s")
		if err == nil {
			return int(d.Seconds()) > getFightDuration(contest)
		}
	}

	if contest.GoldenScore != nil {
		return *contest.GoldenScore == "1"
	}

	return false
}

func getFightDuration(contest judobase.Contest) int {
	if contest.CompYear != nil {
		y, err := strconv.Atoi(*contest.CompYear)
		if err == nil {
			if y <= 2016 && parseGender(contest) == male {
				return 300
			}
		}
	}

	return 240
}
