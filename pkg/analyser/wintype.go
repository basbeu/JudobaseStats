package analyser

import (
	"github.com/basbeu/JudobaseStats/internal/judobase"
)

type winType int

const (
	winByIppon winType = iota
	winByWaza
	winByYuko
	winBy3Shidos
	winByHansokuMake
	winUnknown
)

func (t winType) string() string {
	switch t {
	case winByIppon:
		return "Ippon"
	case winByWaza:
		return "Waza-Ari"
	case winBy3Shidos:
		return "3 shidos"
	case winByHansokuMake:
		return "Hansoku-make"
	default:
		return "Unknown"
	}
}

func parseWinType(contest judobase.Contest) winType {
	if isWinnerWhite(contest) {
		if (contest.IpponWhite != nil && *contest.IpponWhite == "1") || (contest.WazaWhite != nil && *contest.WazaWhite == "2") {
			if contest.PenaltyBlue != nil && *contest.PenaltyBlue == "3" {
				return winBy3Shidos
			} else if contest.HSKBlue != nil && *contest.HSKBlue == "1" {
				return winByHansokuMake
			}
			return winByIppon
		} else if contest.WazaWhite != nil && *contest.WazaWhite == "1" {
			return winByWaza
		} else if contest.YukoWhite != nil && *contest.YukoWhite != "0" {
			return winByYuko
		}
	} else if isWinnerBlue(contest) {
		if (contest.IpponBlue != nil && *contest.IpponBlue == "1") || (contest.WazaBlue != nil && *contest.WazaBlue == "2") {
			if contest.PenaltyWhite != nil && *contest.PenaltyWhite == "3" {
				return winBy3Shidos
			} else if contest.HSKWhite != nil && *contest.HSKWhite == "1" {
				return winByHansokuMake
			}
			return winByIppon
		} else if contest.WazaBlue != nil && *contest.WazaBlue == "1" {
			return winByWaza
		} else if contest.YukoBlue != nil && *contest.YukoBlue != "0" {
			return winByYuko
		}
	}
	return winUnknown
}

func isWinnerWhite(contest judobase.Contest) bool {
	return contest.IDWinner != nil && contest.IDPersonWhite != nil && *contest.IDPersonWhite == *contest.IDWinner
}

func isWinnerBlue(contest judobase.Contest) bool {
	return contest.IDWinner != nil && contest.IDPersonBlue != nil && *contest.IDPersonBlue == *contest.IDWinner
}
