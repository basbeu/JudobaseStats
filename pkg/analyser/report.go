package analyser

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type Reporter interface {
	ReportCategoryStats(competitionName string, categoryName string, winRecords []WinRecord)
	ReportGenderStats(competitionName string, winRecords []WinRecord)
	Close() error
}

func NewReporter(mode string, outputPath string) Reporter {
	switch mode {
	case "stdout":
		return &stdOutReporter{}
	case "txt":
		return &txtReporter{
			outputPath: outputPath,
		}
	case "csv":
		return &csvReporter{
			outputPath: outputPath,
			csvContent: [][]string{
				{"competition", "category", "round", "fights", "ippon", "waza", "shidos", "hsk", "unknown", "regular", "golden"},
			},
		}
	default:
		return &stdOutReporter{}
	}
}

type stdOutReporter struct{}

func (r *stdOutReporter) ReportCategoryStats(competitionName string, categoryName string, winRecords []WinRecord) {
	reportCategoryStats(os.Stdout, competitionName, categoryName, winRecords)
}

func (r *stdOutReporter) ReportGenderStats(competitionName string, winRecords []WinRecord) {
	reportGenderStats(os.Stdout, competitionName, winRecords)
}

func (r *stdOutReporter) Close() error {
	return nil
}

type txtReporter struct {
	outputPath string
}

func (r *txtReporter) ReportCategoryStats(competitionName string, categoryName string, winRecords []WinRecord) {
	file, err := os.OpenFile(fmt.Sprintf("%s/analysis-%s%s.txt", r.outputPath, competitionName, categoryName), os.O_CREATE, 0666)
	if err != nil {
		return
	}
	defer file.Close()
	reportCategoryStats(file, competitionName, categoryName, winRecords)
}

func (r *txtReporter) ReportGenderStats(competitionName string, winRecords []WinRecord) {
	file, err := os.OpenFile(fmt.Sprintf("%s/analysis-%s%s.txt", r.outputPath, competitionName, "gender"), os.O_CREATE, 0666)
	if err != nil {
		return
	}
	defer file.Close()

	reportGenderStats(file, competitionName, winRecords)
}

func (r *txtReporter) Close() error {
	return nil
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

func reportGenderStats(out io.Writer, competitionName string, winRecords []WinRecord) {
	if len(winRecords) > 0 {
		fmt.Fprintln(out, "====================================")
		fmt.Fprintf(out, "Competition: %s\n", competitionName)

		winsByGender := groupByGender(winRecords)
		for _, gender := range genders {
			fmt.Fprintln(out, "====================================")
			fmt.Fprintf(out, "Gender: %s\n", gender.string())
			groupByRound := groupByRound(winsByGender[gender])
			for _, round := range rounds {
				reportRoundStats(out, round.string(), groupByRound[round])
			}
			reportRoundStats(out, "ALL ROUNDS", winsByGender[gender])
		}
	}
}

func reportRoundStats(out io.Writer, round string, winRecords []WinRecord) {
	if len(winRecords) > 0 {
		winsByTypes := groupByWinType(winRecords).count()
		winsByFinishMode := groupByFinishMode(winRecords).count()
		fmt.Fprintln(out, "======== "+round+" ================")
		fmt.Fprintf(out, "# fights: %d\n", len(winRecords))
		fmt.Fprintf(out, "# wins by ippon: %d %s\n", winsByTypes[winByIppon], formatPercentage(winsByTypes[winByIppon], len(winRecords)))
		fmt.Fprintf(out, "# wins by waza-ari: %d %s\n", winsByTypes[winByWaza], formatPercentage(winsByTypes[winByWaza], len(winRecords)))
		fmt.Fprintf(out, "# wins by 3 shidos: %d %s\n", winsByTypes[winByShido], formatPercentage(winsByTypes[winByShido], len(winRecords)))
		fmt.Fprintf(out, "# wins by direct hansoku-make: %d %s\n", winsByTypes[winByHansokuMake], formatPercentage(winsByTypes[winByHansokuMake], len(winRecords)))
		fmt.Fprintf(out, "# unknown win types: %d %s\n", winsByTypes[winUnknown], formatPercentage(winsByTypes[winUnknown], len(winRecords)))
		fmt.Fprintln(out, "------------------------------------")
		fmt.Fprintf(out, "# wins in regular time: %d %s\n", winsByFinishMode[regularTime], formatPercentage(winsByFinishMode[regularTime], len(winRecords)))
		fmt.Fprintf(out, "# wins in Golden Score: %d %s\n", winsByFinishMode[goldenScore], formatPercentage(winsByFinishMode[goldenScore], len(winRecords)))
		fmt.Fprintln(out, "====================================")
	}
}

func formatPercentage(part int, total int) string {
	percentage := float64(part) / float64(total) * 100
	return fmt.Sprintf("(%.2f%%)", percentage)
}

type csvReporter struct {
	outputPath string
	csvContent [][]string
}

func (r *csvReporter) ReportCategoryStats(competitionName string, categoryName string, winRecords []WinRecord) {
	groupByRound := groupByRound(winRecords)
	for _, round := range rounds {
		r.csvContent = append(r.csvContent, reportRoundStatsArray(competitionName, categoryName, round.string(), groupByRound[round]))
	}
	r.csvContent = append(r.csvContent, reportRoundStatsArray(competitionName, categoryName, "all", winRecords))
}

func (r *csvReporter) ReportGenderStats(competitionName string, winRecords []WinRecord) {
	if len(winRecords) > 0 {
		winsByGender := groupByGender(winRecords)

		for _, gender := range genders {
			groupByRound := groupByRound(winsByGender[gender])
			for _, round := range rounds {
				r.csvContent = append(r.csvContent, reportRoundStatsArray(competitionName, gender.string(), round.string(), groupByRound[round]))
			}

			r.csvContent = append(r.csvContent, reportRoundStatsArray(competitionName, gender.string(), "all", winsByGender[gender]))
		}
	}
}

func (r *csvReporter) Close() error {
	if len(r.csvContent) >= 1 && len(r.csvContent[0]) >= 0 {
		file, err := os.OpenFile(fmt.Sprintf("%s/analysis-%s.csv", r.outputPath, r.csvContent[1][0]), os.O_CREATE, 0666)
		if err != nil {
			return err
		}
		defer file.Close()
		w := csv.NewWriter(file)
		return w.WriteAll(r.csvContent)
	}
	return nil
}

func reportRoundStatsArray(competitionName string, categoryName string, round string, winRecords []WinRecord) []string {
	winsByTypes := groupByWinType(winRecords).count()
	winsByFinishMode := groupByFinishMode(winRecords).count()
	return []string{
		competitionName,
		categoryName,
		round,
		fmt.Sprint(len(winRecords)),
		fmt.Sprint(winsByTypes[winByIppon]),
		fmt.Sprint(winsByTypes[winByWaza]),
		fmt.Sprint(winsByTypes[winByShido]),
		fmt.Sprint(winsByTypes[winByHansokuMake]),
		fmt.Sprint(winsByTypes[winUnknown]),
		fmt.Sprint(winsByFinishMode[regularTime]),
		fmt.Sprint(winsByFinishMode[goldenScore]),
	}
}
