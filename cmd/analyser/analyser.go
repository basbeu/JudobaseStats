package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/basbeu/JudobaseStats/internal/judobase"
	"github.com/basbeu/JudobaseStats/pkg/analyser"
)

func main() {
	competitionID := flag.String("competition", "2653", "competition ID in judobase.ijf.org")
	inputPath := flag.String("input", "../../data", "path of the input folder")
	outputMode := flag.String("outMode", "stdout", "output mode")
	outputPath := flag.String("output", "../../analysis", "path of the output folder")

	flag.Parse()

	folderInputPath := fmt.Sprintf("%s/%s", *inputPath, *competitionID)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	competition, err := judobase.ParseCompetition(folderInputPath, logger)
	if err != nil {
		logger.Error("failed to parse competition", "err", err)
		return
	}

	reporter := analyser.NewReporter(*outputMode, *outputPath)

	overallWinRecords := []analyser.WinRecord{}
	for _, category := range competition.Categories {
		winRecords := analyser.ParseWinRecords(category.Contests)
		reporter.ReportCategoryStats(competition.Name, category.Name, winRecords)
		overallWinRecords = append(overallWinRecords, winRecords...)
	}
	reporter.ReportCategoryStats(competition.Name, "all", overallWinRecords)
	reporter.ReportGenderStats(competition.Name, overallWinRecords)
}
