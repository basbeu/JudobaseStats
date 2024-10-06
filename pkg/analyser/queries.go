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
	winRecordsByType := groupByResult{
		winByIppon:       {},
		winByWaza:        {},
		winBy3Shidos:     {},
		winByHansokuMake: {},
		winUnknown:       {},
	}
	for _, r := range winRecords {
		winRecordsByType[r.winType] = append(winRecordsByType[r.winType], r)
	}
	return winRecordsByType
}

func groupByFinishMode(winRecords []WinRecord) groupByResult {
	winRecordsByFinishMode := groupByResult{
		goldenScore: {},
		regularTime: {},
	}
	for _, r := range winRecords {
		winRecordsByFinishMode[r.finishMode] = append(winRecordsByFinishMode[r.finishMode], r)
	}

	return winRecordsByFinishMode
}

func groupByRound(winRecords []WinRecord) groupByResult {
	winRecordsByRound := groupByResult{
		round64:      {},
		round32:      {},
		round16:      {},
		quarterFinal: {},
		semiFinal:    {},
		repechage:    {},
		bronze:       {},
		final:        {},
		unknownRound: {},
	}
	for _, r := range winRecords {
		winRecordsByRound[r.round] = append(winRecordsByRound[r.round], r)
	}
	return winRecordsByRound
}

func groupByGender(winRecords []WinRecord) groupByResult {
	winRecordsByGender := groupByResult{
		male:          {},
		female:        {},
		unknownGender: {},
	}
	for _, r := range winRecords {
		winRecordsByGender[r.gender] = append(winRecordsByGender[r.gender], r)
	}
	return winRecordsByGender
}
