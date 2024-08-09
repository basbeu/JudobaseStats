package analyser

import "github.com/basbeu/JudobaseStats/internal/judobase"

type round int

const (
	round64 round = iota
	round32
	round16
	quarterFinal
	semiFinal
	repechage
	bronze
	final
	unknown
)

var rounds []round

func init() {
	rounds = []round{
		round64, round32, round16, quarterFinal, semiFinal, repechage, bronze, final,
	}
}

func (r round) string() string {
	switch r {
	case round64:
		return "Round of 64"
	case round32:
		return "Round of 32"
	case round16:
		return "Round of 16"
	case quarterFinal:
		return "Quarter-Final"
	case semiFinal:
		return "Semi-Final"
	case repechage:
		return "Repechage"
	case bronze:
		return "Bronze"
	case final:
		return "Final"
	default:
		return "Unknown"
	}
}

func parseRound(contest judobase.Contest) round {
	if contest.RoundName != nil {
		switch *contest.RoundName {
		case "Round of 64":
			return round64
		case "Round of 32":
			return round32
		case "Round of 16":
			return round16
		case "Quarter-Final":
			return quarterFinal
		case "Semi-Final":
			return semiFinal
		case "Repechage":
			return repechage
		case "Bronze":
			return bronze
		case "Final":
			return final
		default:
			return unknown
		}
	}
	return unknown
}
