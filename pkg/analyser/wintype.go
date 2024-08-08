package analyser

import "github.com/basbeu/JudobaseStats/internal/judobase"

type WinType int

const (
	WinByIppon WinType = iota
	WinByWaza
	WinByShido
	WinByHansokuMake
	WinUnknown
)

func (t WinType) String() string {
	switch t {
	case WinByIppon:
		return "Ippon"
	case WinByWaza:
		return "Waza-Ari"
	case WinByShido:
		return "3 shidos"
	case WinByHansokuMake:
		return "Hansoku-make"
	default:
		return "Unknown"
	}
}

func parseWinType(contest judobase.Contest) WinType {
	if contest.IDWinner != nil && contest.IDPersonWhite != nil && *contest.IDPersonWhite == *contest.IDWinner {
		if (contest.IpponWhite != nil && *contest.IpponWhite == "1") || (contest.WazaWhite != nil && *contest.WazaWhite == "2") {
			if contest.PenaltyBlue != nil && *contest.PenaltyBlue == "3" {
				return WinByShido
			} else if contest.HSKBlue != nil && *contest.HSKBlue == "1" {
				return WinByHansokuMake
			}
			return WinByIppon
		} else if contest.WazaWhite != nil && *contest.WazaWhite == "1" {
			return WinByWaza
		}
	} else if contest.IDWinner != nil && contest.IDPersonBlue != nil && *contest.IDPersonBlue == *contest.IDWinner {
		if (contest.IpponBlue != nil && *contest.IpponBlue == "1") || (contest.WazaBlue != nil && *contest.WazaBlue == "2") {
			if contest.PenaltyWhite != nil && *contest.PenaltyWhite == "3" {
				return WinByShido
			} else if contest.HSKWhite != nil && *contest.HSKWhite == "1" {
				return WinByHansokuMake
			}
			return WinByIppon
		} else if contest.WazaBlue != nil && *contest.WazaBlue == "1" {
			return WinByWaza
		}
	}
	return WinUnknown
}
