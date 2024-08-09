package analyser

type groupByWinResult map[WinType][]WinRecord

func (g groupByWinResult) count() map[WinType]int {
	return map[WinType]int{
		WinByIppon:       len(g[WinByIppon]),
		WinByWaza:        len(g[WinByWaza]),
		WinByShido:       len(g[WinByShido]),
		WinByHansokuMake: len(g[WinByHansokuMake]),
		WinUnknown:       len(g[WinUnknown]),
	}
}

func groupByWinType(winRecords []WinRecord) groupByWinResult {
	winRecordByType := groupByWinResult{
		WinByIppon:       {},
		WinByWaza:        {},
		WinByShido:       {},
		WinByHansokuMake: {},
		WinUnknown:       {},
	}
	for _, r := range winRecords {
		winRecordByType[r.Type] = append(winRecordByType[r.Type], r)
	}
	return winRecordByType
}

type groupByGoldenScoreResult map[bool][]WinRecord

func (g groupByGoldenScoreResult) count() map[bool]int {
	return map[bool]int{
		true:  len(g[true]),
		false: len(g[false]),
	}
}

func groupByGoldenScore(winRecords []WinRecord) groupByGoldenScoreResult {
	winByGolden := groupByGoldenScoreResult{
		true:  {},
		false: {},
	}
	for _, r := range winRecords {
		winByGolden[r.GoldenScore] = append(winByGolden[r.GoldenScore], r)
	}

	return winByGolden
}

type groupByRoundResult map[Round][]WinRecord

func groupByRound(winRecords []WinRecord) groupByRoundResult {
	winRecordByRound := groupByRoundResult{
		Round64:      {},
		Round32:      {},
		Round16:      {},
		QuarterFinal: {},
		SemiFinal:    {},
		Repechage:    {},
		Bronze:       {},
		Final:        {},
		Unknown:      {},
	}
	for _, r := range winRecords {
		winRecordByRound[r.Round] = append(winRecordByRound[r.Round], r)
	}
	return winRecordByRound
}
