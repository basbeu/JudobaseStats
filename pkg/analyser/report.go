package analyser

func GroupByWinType(winRecords []WinRecord) map[WinType]int {
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

func GroupByGoldenScore(winRecords []WinRecord) map[bool]int {
	winByGolden := map[bool]int{
		true:  0,
		false: 0,
	}
	for _, r := range winRecords {
		winByGolden[r.GoldenScore]++
	}

	return winByGolden
}
