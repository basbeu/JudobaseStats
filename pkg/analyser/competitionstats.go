package analyser

type CompetitionStats struct {
	Name  string
	Stats []CategoryStats
}

type CategoryStats struct {
	Category string
	Stats    []RoundStats
}

type RoundStats struct {
	Round       string
	Fights      int
	Ippon       int
	Waza        int
	Shidos      int
	Hansokumake int
	Unknown     int
	Regular     int
	Golden      int
}

func newRoundStats(round string, winRecords []WinRecord) RoundStats {
	winsByTypes := groupByWinType(winRecords).count()
	winsByFinishMode := groupByFinishMode(winRecords).count()

	return RoundStats{
		Round:       round,
		Fights:      len(winRecords),
		Ippon:       winsByTypes[winByIppon],
		Waza:        winsByTypes[winByWaza],
		Shidos:      winsByTypes[winByShido],
		Hansokumake: winsByTypes[winByHansokuMake],
		Unknown:     winsByTypes[winUnknown],
		Regular:     winsByFinishMode[regularTime],
		Golden:      winsByFinishMode[goldenScore],
	}
}
