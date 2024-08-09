package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/basbeu/JudobaseStats/internal/category"
	"github.com/basbeu/JudobaseStats/pkg/scraper"
)

func main() {
	competitionID := flag.String("competition", "2653", "competition ID in judobase.ijf.org")
	outputPath := flag.String("output", "../../data", "path of the output folder")

	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	folderOutputPath := fmt.Sprintf("%s/%s", *outputPath, *competitionID)
	err := os.MkdirAll(folderOutputPath, 0700)
	if err != nil {
		logger.Error("failed to create directory", "err", err)
		return
	}

	scraperClient := scraper.NewScraperClient(*competitionID, logger)

	for _, category := range category.Categories {
		categoryBytes, err := scraperClient.ScrapeCategory(category)
		if err != nil {
			logger.Error("failed to scrape", "err", err)
			return
		}

		filename := fmt.Sprintf("%s/%s.json", folderOutputPath, category.String())
		os.WriteFile(filename, categoryBytes, 0666)
	}
}
