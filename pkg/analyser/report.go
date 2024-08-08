package analyser

import "fmt"

func DisplayWinRecords(competitionName string, categoryName string, winRecords []WinRecord) {
	winByTypes := groupByWinType(winRecords)
	winByGolden := groupByGoldenScore(winRecords)

	fmt.Println("====================================")
	fmt.Printf("Competition: %s\n", competitionName)
	fmt.Printf("Category: %s\n", categoryName)
	fmt.Printf("# fights: %d\n", len(winRecords))
	fmt.Println("====================================")
	fmt.Printf("# wins by ippon: %d %s\n", winByTypes[WinByIppon], formatPercentage(winByTypes[WinByIppon], len(winRecords)))
	fmt.Printf("# wins by waza-ari: %d %s\n", winByTypes[WinByWaza], formatPercentage(winByTypes[WinByWaza], len(winRecords)))
	fmt.Printf("# wins by 3 shidos: %d %s\n", winByTypes[WinByShido], formatPercentage(winByTypes[WinByShido], len(winRecords)))
	fmt.Printf("# wins by direct hansoku-make: %d %s\n", winByTypes[WinByHansokuMake], formatPercentage(winByTypes[WinByHansokuMake], len(winRecords)))
	fmt.Println("====================================")
	fmt.Printf("# wins in regular time: %d %s\n", winByGolden[false], formatPercentage(winByGolden[false], len(winRecords)))
	fmt.Printf("# wins in Golden Score: %d %s\n", winByGolden[true], formatPercentage(winByGolden[true], len(winRecords)))
	fmt.Println("====================================")
}

func groupByWinType(winRecords []WinRecord) map[WinType]int {
	winByTypes := map[WinType]int{
		WinByIppon:       0,
		WinByWaza:        0,
		WinByShido:       0,
		WinByHansokuMake: 0,
		WinUnknown:       0,
	}
	for _, r := range winRecords {
		winByTypes[r.Type]++
	}
	return winByTypes
}

func groupByGoldenScore(winRecords []WinRecord) map[bool]int {
	winByGolden := map[bool]int{
		true:  0,
		false: 0,
	}
	for _, r := range winRecords {
		winByGolden[r.GoldenScore]++
	}

	return winByGolden
}

func formatPercentage(part int, total int) string {
	percentage := float64(part) / float64(total) * 100
	return fmt.Sprintf("(%.2f%%)", percentage)
}
