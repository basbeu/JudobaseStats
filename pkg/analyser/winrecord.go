package analyser

import "github.com/basbeu/JudobaseStats/internal/judobase"

type WinRecord struct {
	winType    winType
	finishMode finishMode
	round      round
}

func newWinRecord(contest judobase.Contest) WinRecord {
	return WinRecord{
		winType:    parseWinType(contest),
		finishMode: parseFinishMode(contest),
		round:      parseRound(contest),
	}
}

func ParseWinRecords(contests []judobase.Contest) []WinRecord {
	winRecords := []WinRecord{}
	for _, contest := range contests {
		winRecords = append(winRecords, newWinRecord(contest))
	}

	return winRecords
}
