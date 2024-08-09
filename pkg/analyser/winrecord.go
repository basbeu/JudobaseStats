package analyser

import "github.com/basbeu/JudobaseStats/internal/judobase"

type WinRecord struct {
	Type       WinType
	FinishMode FinishMode
	Round      Round
}

func newWinRecord(contest judobase.Contest) WinRecord {
	return WinRecord{
		Type:       parseWinType(contest),
		FinishMode: parseFinishMode(contest),
		Round:      parseRound(contest),
	}
}

func ParseWinRecords(contests []judobase.Contest) []WinRecord {
	winRecords := []WinRecord{}
	for _, contest := range contests {
		winRecords = append(winRecords, newWinRecord(contest))
	}

	return winRecords
}
