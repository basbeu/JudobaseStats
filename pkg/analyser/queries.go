package analyser

type groupingKey interface {
	string() string
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
		winByIppon:       {},
		winByWaza:        {},
		winByShido:       {},
		winByHansokuMake: {},
		winUnknown:       {},
	}
	for _, r := range winRecords {
		winRecordByType[r.winType] = append(winRecordByType[r.winType], r)
	}
	return winRecordByType
}

func groupByFinishMode(winRecords []WinRecord) groupByResult {
	winByFinishMode := groupByResult{
		goldenScore: {},
		regularTime: {},
	}
	for _, r := range winRecords {
		winByFinishMode[r.finishMode] = append(winByFinishMode[r.finishMode], r)
	}

	return winByFinishMode
}

func groupByRound(winRecords []WinRecord) groupByResult {
	winRecordByRound := groupByResult{
		round64:      {},
		round32:      {},
		round16:      {},
		quarterFinal: {},
		semiFinal:    {},
		repechage:    {},
		bronze:       {},
		final:        {},
		unknown:      {},
	}
	for _, r := range winRecords {
		winRecordByRound[r.round] = append(winRecordByRound[r.round], r)
	}
	return winRecordByRound
}
