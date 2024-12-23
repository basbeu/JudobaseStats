package analyser

import (
	"strconv"

	"github.com/basbeu/JudobaseStats/internal/judobase"
)

type winType int

const (
	winByIppon winType = iota
	winByWaza
	winByYuko
	winByMaxShidos
	winByHansokuMake
	winByShido
	winUnknown
)

func (t winType) string() string {
	switch t {
	case winByIppon:
		return "Ippon"
	case winByWaza:
		return "Waza-Ari"
	case winByMaxShidos:
		return "max shidos"
	case winByShido:
		return "win by shido"
	case winByHansokuMake:
		return "Hansoku-make"
	default:
		return "Unknown"
	}
}

func parseWinType(contest judobase.Contest) winType {
	if isWinnerWhite(contest) {
		if (contest.IpponWhite != nil && *contest.IpponWhite == "1") || (contest.WazaWhite != nil && *contest.WazaWhite == "2") {
			if contest.PenaltyBlue != nil && *contest.PenaltyBlue == getMaxShidos(contest) {
				return winByMaxShidos
			} else if contest.HSKBlue != nil && *contest.HSKBlue == "1" {
				return winByHansokuMake
			}
			return winByIppon
		} else if contest.WazaWhite != nil && *contest.WazaWhite == "1" {
			return winByWaza
		} else if contest.YukoWhite != nil && *contest.YukoWhite != "0" {
			return winByYuko
		} else if *contest.PenaltyBlue > *contest.PenaltyWhite {
			return winByShido
		}
	} else if isWinnerBlue(contest) {
		if (contest.IpponBlue != nil && *contest.IpponBlue == "1") || (contest.WazaBlue != nil && *contest.WazaBlue == "2") {
			if contest.PenaltyWhite != nil && *contest.PenaltyWhite == getMaxShidos(contest) {
				return winByMaxShidos
			} else if contest.HSKWhite != nil && *contest.HSKWhite == "1" {
				return winByHansokuMake
			}
			return winByIppon
		} else if contest.WazaBlue != nil && *contest.WazaBlue == "1" {
			return winByWaza
		} else if contest.YukoBlue != nil && *contest.YukoBlue != "0" {
			return winByYuko
		} else if *contest.PenaltyWhite > *contest.PenaltyBlue {
			return winByShido
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

func getMaxShidos(contest judobase.Contest) string {
	if contest.CompYear != nil {
		y, err := strconv.Atoi(*contest.CompYear)
		if err == nil {
			if y <= 2016 {
				return "4"
			}
		}
	}
	return "3"
}
