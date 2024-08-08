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

	flag.Parse()

	folderInputPath := fmt.Sprintf("%s/%s", *inputPath, *competitionID)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	competition, err := judobase.ParseCompetition(folderInputPath, logger)
	if err != nil {
		logger.Error("failed to parse competition", "err", err)
		return
	}

	for _, category := range competition.Categories {
		winRecords := analyser.ParseWinRecords(category.Contests)
		winByTypes := analyser.GroupByWinType(winRecords)
		winByGolden := analyser.GroupByGoldenScore(winRecords)

		fmt.Println("====================================")
		fmt.Printf("Competition: %s\n", competition.Name)
		fmt.Printf("Category: %s\n", category.Name)
		fmt.Printf("# fights: %d\n", len(winRecords))
		fmt.Println("====================================")
		fmt.Printf("# wins by ippon: %d %s\n", winByTypes[analyser.WinByIppon], formatPercentage(winByTypes[analyser.WinByIppon], len(winRecords)))
		fmt.Printf("# wins by waza-ari: %d %s\n", winByTypes[analyser.WinByWaza], formatPercentage(winByTypes[analyser.WinByWaza], len(winRecords)))
		fmt.Printf("# wins by 3 shidos: %d %s\n", winByTypes[analyser.WinByShido], formatPercentage(winByTypes[analyser.WinByShido], len(winRecords)))
		fmt.Printf("# wins by direct hansoku-make: %d %s\n", winByTypes[analyser.WinByHansokuMake], formatPercentage(winByTypes[analyser.WinByHansokuMake], len(winRecords)))
		fmt.Println("====================================")
		fmt.Printf("# wins in regular time: %d %s\n", winByGolden[false], formatPercentage(winByGolden[false], len(winRecords)))
		fmt.Printf("# wins in Golden Score: %d %s\n", winByGolden[true], formatPercentage(winByGolden[true], len(winRecords)))
		fmt.Println("====================================")
	}
}

func formatPercentage(part int, total int) string {
	percentage := float64(part) / float64(total) * 100
	return fmt.Sprintf("(%.2f%%)", percentage)
}
