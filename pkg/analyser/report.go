package analyser

import (
	"fmt"
	"io"
	"os"
)

type Reporter interface {
	ReportCategoryStats(competitionName string, categoryName string, winRecords []WinRecord)
}

func NewReporter(mode string, outputPath string) Reporter {
	switch mode {
	case "stdout":
		return stdOutReporter{}
	case "txt":
		return txtReporter{
			outputPath: outputPath,
		}
	default:
		return stdOutReporter{}
	}
}

type stdOutReporter struct{}

func (r stdOutReporter) ReportCategoryStats(competitionName string, categoryName string, winRecords []WinRecord) {
	reportCategoryStats(os.Stdout, competitionName, categoryName, winRecords)
}

type txtReporter struct {
	outputPath string
}

func (r txtReporter) ReportCategoryStats(competitionName string, categoryName string, winRecords []WinRecord) {
	file, err := os.OpenFile(fmt.Sprintf("%s/analysis-%s%s.txt", r.outputPath, competitionName, categoryName), os.O_CREATE, 0666)
	if err != nil {
		return
	}
	defer file.Close()
	reportCategoryStats(file, competitionName, categoryName, winRecords)
}

func reportCategoryStats(out io.Writer, competitionName string, categoryName string, winRecords []WinRecord) {
	fmt.Fprintln(out, "====================================")
	fmt.Fprintf(out, "Competition: %s\n", competitionName)
	fmt.Fprintf(out, "Category: %s\n", categoryName)
	fmt.Fprintf(out, "# fights: %d\n", len(winRecords))

	groupByRound := groupByRound(winRecords)
	for _, round := range rounds {
		reportRoundStats(out, round.string(), groupByRound[round])
	}
	reportRoundStats(out, "ALL ROUNDS", winRecords)
}

func reportRoundStats(out io.Writer, round string, winRecords []WinRecord) {
	if len(winRecords) > 0 {
		winByTypes := groupByWinType(winRecords).count()
		winByFinishMode := groupByFinishMode(winRecords).count()
		fmt.Fprintln(out, "======== "+round+" ================")
		fmt.Fprintf(out, "# fights: %d\n", len(winRecords))
		fmt.Fprintf(out, "# wins by ippon: %d %s\n", winByTypes[winByIppon], formatPercentage(winByTypes[winByIppon], len(winRecords)))
		fmt.Fprintf(out, "# wins by waza-ari: %d %s\n", winByTypes[winByWaza], formatPercentage(winByTypes[winByWaza], len(winRecords)))
		fmt.Fprintf(out, "# wins by 3 shidos: %d %s\n", winByTypes[winByShido], formatPercentage(winByTypes[winByShido], len(winRecords)))
		fmt.Fprintf(out, "# wins by direct hansoku-make: %d %s\n", winByTypes[winByHansokuMake], formatPercentage(winByTypes[winByHansokuMake], len(winRecords)))
		fmt.Fprintln(out, "------------------------------------")
		fmt.Fprintf(out, "# wins in regular time: %d %s\n", winByFinishMode[regularTime], formatPercentage(winByFinishMode[regularTime], len(winRecords)))
		fmt.Fprintf(out, "# wins in Golden Score: %d %s\n", winByFinishMode[goldenScore], formatPercentage(winByFinishMode[goldenScore], len(winRecords)))
		fmt.Fprintln(out, "====================================")
	}
}

func formatPercentage(part int, total int) string {
	percentage := float64(part) / float64(total) * 100
	return fmt.Sprintf("(%.2f%%)", percentage)
}
