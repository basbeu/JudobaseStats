package analyser

import "github.com/basbeu/JudobaseStats/internal/judobase"

type Round int

const (
	Round64 Round = iota
	Round32
	Round16
	QuarterFinal
	SemiFinal
	Repechage
	Bronze
	Final
	Unknown
)

var Rounds []Round

func init() {
	Rounds = []Round{
		Round64, Round32, Round16, QuarterFinal, SemiFinal, Repechage, Bronze, Final,
	}
}

func (r Round) String() string {
	switch r {
	case Round64:
		return "Round of 64"
	case Round32:
		return "Round of 32"
	case Round16:
		return "Round of 16"
	case QuarterFinal:
		return "Quarter-Final"
	case SemiFinal:
		return "Semi-Final"
	case Repechage:
		return "Repechage"
	case Bronze:
		return "Bronze"
	case Final:
		return "Final"
	default:
		return "Unknown"
	}
}

func parseRound(contest judobase.Contest) Round {
	if contest.RoundName != nil {
		switch *contest.RoundName {
		case "Round of 64":
			return Round64
		case "Round of 32":
			return Round32
		case "Round of 16":
			return Round16
		case "Quarter-Final":
			return QuarterFinal
		case "Semi-Final":
			return SemiFinal
		case "Repechage":
			return Repechage
		case "Bronze":
			return Bronze
		case "Final":
			return Final
		default:
			return Unknown
		}
	}
	return Unknown
}
