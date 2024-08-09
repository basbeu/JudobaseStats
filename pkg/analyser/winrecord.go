package analyser

import "github.com/basbeu/JudobaseStats/internal/judobase"

type WinRecord struct {
	Type        WinType
	GoldenScore bool
	Round       Round
}

func newWinRecord(contest judobase.Contest) WinRecord {
	return WinRecord{
		Type:        parseWinType(contest),
		GoldenScore: parseGoldenScore(contest),
		Round:       parseRound(contest),
	}
}

func parseGoldenScore(contest judobase.Contest) bool {
	return contest.GoldenScore != nil && *contest.GoldenScore == "1"
}

func ParseWinRecords(contests []judobase.Contest) []WinRecord {
	winRecords := []WinRecord{}
	for _, contest := range contests {
		winRecords = append(winRecords, newWinRecord(contest))
	}

	return winRecords
}
