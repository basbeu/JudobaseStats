package analyser

import (
	"github.com/basbeu/JudobaseStats/internal/judobase"
)

type winType int

const (
	winByIppon winType = iota
	winByWaza
	winByShido
	winByHansokuMake
	winUnknown
)

func (t winType) string() string {
	switch t {
	case winByIppon:
		return "Ippon"
	case winByWaza:
		return "Waza-Ari"
	case winByShido:
		return "3 shidos"
	case winByHansokuMake:
		return "Hansoku-make"
	default:
		return "Unknown"
	}
}

func parseWinType(contest judobase.Contest) winType {
	if contest.IDWinner != nil && contest.IDPersonWhite != nil && *contest.IDPersonWhite == *contest.IDWinner {
		if (contest.IpponWhite != nil && *contest.IpponWhite == "1") || (contest.WazaWhite != nil && *contest.WazaWhite == "2") {
			if contest.PenaltyBlue != nil && *contest.PenaltyBlue == "3" {
				return winByShido
			} else if contest.HSKBlue != nil && *contest.HSKBlue == "1" {
				return winByHansokuMake
			}
			return winByIppon
		} else if contest.WazaWhite != nil && *contest.WazaWhite == "1" {
			return winByWaza
		}
	} else if contest.IDWinner != nil && contest.IDPersonBlue != nil && *contest.IDPersonBlue == *contest.IDWinner {
		if (contest.IpponBlue != nil && *contest.IpponBlue == "1") || (contest.WazaBlue != nil && *contest.WazaBlue == "2") {
			if contest.PenaltyWhite != nil && *contest.PenaltyWhite == "3" {
				return winByShido
			} else if contest.HSKWhite != nil && *contest.HSKWhite == "1" {
				return winByHansokuMake
			}
			return winByIppon
		} else if contest.WazaBlue != nil && *contest.WazaBlue == "1" {
			return winByWaza
		}
	}
	return winUnknown
}
