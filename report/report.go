package report

import "github.com/basbeu/JudobaseStats/stats"

func GroupByWinType(winRecords []stats.WinRecord) map[stats.WinType]int {
	winByTypes := map[stats.WinType]int{
		stats.WinByIppon:       0,
		stats.WinByWaza:        0,
		stats.WinByShido:       0,
		stats.WinByHansokuMake: 0,
		stats.WinUnknown:       0,
	}
	for _, r := range winRecords {
		winByTypes[r.Type]++
	}
	return winByTypes
}

func GroupByGoldenScore(winRecords []stats.WinRecord) map[bool]int {
	winByGolden := map[bool]int{
		true:  0,
		false: 0,
	}
	for _, r := range winRecords {
		winByGolden[r.GoldenScore]++
	}

	return winByGolden
}
