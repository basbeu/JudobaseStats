package analyser

type groupingKey interface {
	String() string
}

type groupByResult map[groupingKey][]WinRecord

func (g groupByResult) count() map[groupingKey]int {
	res := map[groupingKey]int{}

	for key, winRecords := range g {
		res[key] = len(winRecords)
	}
	return res
}

func groupByWinType(winRecords []WinRecord) groupByResult {
	winRecordByType := groupByResult{
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

func groupByFinishMode(winRecords []WinRecord) groupByResult {
	winByFinishMode := groupByResult{
		goldenScore: {},
		regularTime: {},
	}
	for _, r := range winRecords {
		winByFinishMode[r.FinishMode] = append(winByFinishMode[r.FinishMode], r)
	}

	return winByFinishMode
}

func groupByRound(winRecords []WinRecord) groupByResult {
	winRecordByRound := groupByResult{
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
