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
	Yuko        int
	Shido       int
	MaxShidos   int
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
		Yuko:        winsByTypes[winByYuko],
		Shido:       winsByTypes[winByShido],
		MaxShidos:   winsByTypes[winByMaxShidos],
		Hansokumake: winsByTypes[winByHansokuMake],
		Unknown:     winsByTypes[winUnknown],
		Regular:     winsByFinishMode[regularTime],
		Golden:      winsByFinishMode[goldenScore],
	}
}
