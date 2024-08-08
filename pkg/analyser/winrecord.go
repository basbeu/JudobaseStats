package analyser

import "github.com/basbeu/JudobaseStats/internal/judobase"

type WinRecord struct {
	Type        WinType
	GoldenScore bool
}

func newWinRecord(contest judobase.Contest) WinRecord {
	return WinRecord{
		Type:        parseWinType(contest),
		GoldenScore: parseGoldenScore(contest),
	}
}

func parseGoldenScore(contest judobase.Contest) bool {
	return contest.GoldenScore != nil && *contest.GoldenScore == "1"
}

func ParseWinRecords(competition judobase.Competition) []WinRecord {
	winRecords := []WinRecord{}
	for _, contest := range competition.Contests {
		winRecords = append(winRecords, newWinRecord(contest))
	}

	return winRecords
}
