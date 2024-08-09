package analyser

import "fmt"

func DisplayCategoryStats(competitionName string, categoryName string, winRecords []WinRecord) {
	fmt.Println("====================================")
	fmt.Printf("Competition: %s\n", competitionName)
	fmt.Printf("Category: %s\n", categoryName)
	fmt.Printf("# fights: %d\n", len(winRecords))

	groupByRound := groupByRound(winRecords)
	for _, round := range rounds {
		displayRoundStats(round.string(), groupByRound[round])
	}
	displayRoundStats("ALL ROUNDS", winRecords)
}

func displayRoundStats(round string, winRecords []WinRecord) {
	if len(winRecords) > 0 {
		winByTypes := groupByWinType(winRecords).count()
		winByFinishMode := groupByFinishMode(winRecords).count()
		fmt.Println("======== " + round + " ================")
		fmt.Printf("# fights: %d\n", len(winRecords))
		fmt.Printf("# wins by ippon: %d %s\n", winByTypes[winByIppon], formatPercentage(winByTypes[winByIppon], len(winRecords)))
		fmt.Printf("# wins by waza-ari: %d %s\n", winByTypes[winByWaza], formatPercentage(winByTypes[winByWaza], len(winRecords)))
		fmt.Printf("# wins by 3 shidos: %d %s\n", winByTypes[winByShido], formatPercentage(winByTypes[winByShido], len(winRecords)))
		fmt.Printf("# wins by direct hansoku-make: %d %s\n", winByTypes[winByHansokuMake], formatPercentage(winByTypes[winByHansokuMake], len(winRecords)))
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
