package analyser

import "fmt"

func DisplayCategoryStats(competitionName string, categoryName string, winRecords []WinRecord) {
	fmt.Println("====================================")
	fmt.Printf("Competition: %s\n", competitionName)
	fmt.Printf("Category: %s\n", categoryName)
	fmt.Printf("# fights: %d\n", len(winRecords))

	groupByRound := groupByRound(winRecords)
	for _, round := range Rounds {
		displayRoundStats(round.String(), groupByRound[round])
	}
	displayRoundStats("ALL ROUNDS", winRecords)
}

func displayRoundStats(round string, winRecords []WinRecord) {
	if len(winRecords) > 0 {
		winByTypes := groupByWinType(winRecords).count()
		winByFinishMode := groupByFinishMode(winRecords).count()
		fmt.Println("======== " + round + " ================")
		fmt.Printf("# fights: %d\n", len(winRecords))
		fmt.Printf("# wins by ippon: %d %s\n", winByTypes[WinByIppon], formatPercentage(winByTypes[WinByIppon], len(winRecords)))
		fmt.Printf("# wins by waza-ari: %d %s\n", winByTypes[WinByWaza], formatPercentage(winByTypes[WinByWaza], len(winRecords)))
		fmt.Printf("# wins by 3 shidos: %d %s\n", winByTypes[WinByShido], formatPercentage(winByTypes[WinByShido], len(winRecords)))
		fmt.Printf("# wins by direct hansoku-make: %d %s\n", winByTypes[WinByHansokuMake], formatPercentage(winByTypes[WinByHansokuMake], len(winRecords)))
		fmt.Println("------------------------------------")
		fmt.Printf("# wins in regular time: %d %s\n", winByFinishMode[regularTime], formatPercentage(winByFinishMode[regularTime], len(winRecords)))
		fmt.Printf("# wins in Golden Score: %d %s\n", winByFinishMode[goldenScore], formatPercentage(winByFinishMode[goldenScore], len(winRecords)))
		fmt.Println("====================================")
	}
}

func formatPercentage(part int, total int) string {
	percentage := float64(part) / float64(total) * 100
	return fmt.Sprintf("(%.2f%%)", percentage)
}
