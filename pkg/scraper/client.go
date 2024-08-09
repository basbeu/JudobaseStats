package scraper

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/basbeu/JudobaseStats/internal/category"
)

type ScraperClient struct {
	competitionID string
	logger        *slog.Logger
}

func NewScraperClient(competitionID string, logger *slog.Logger) *ScraperClient {
	return &ScraperClient{
		competitionID: competitionID,
		logger:        logger,
	}
}

func (s *ScraperClient) ScrapeCategory(category category.Category) ([]byte, error) {
	url := fmt.Sprintf("https://data.ijf.org/api/get_json?params[action]=contest.find&params[id_competition]=%s&params[id_weight]=%d&params[order_by]=cnum", s.competitionID, category.WeightID())

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		s.logger.Warn("failed to create request", "err", err)
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		s.logger.Warn("failed to query server", "err", err)
		return nil, err
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		s.logger.Warn("failed to read the body", "err", err)
		return nil, err
	}

	return bodyBytes, nil
}
