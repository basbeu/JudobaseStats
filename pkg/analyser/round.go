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
	unknownRound
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
		switch *contest.Round {
		case "5":
			return round64
		case "4":
			return round32
		case "3":
			return round16
		case "2":
			return quarterFinal
		case "1":
			if *contest.RoundName == "Semi-Final" {
				return semiFinal
			}
			return repechage
		case "0":
			if *contest.RoundName == "Final" {
				return final
			}
			return bronze
		default:
			return unknownRound
		}
	}
	return unknownRound
}
