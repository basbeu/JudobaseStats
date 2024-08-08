package judobase

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"os"
)

func ParseCompetition(folderInputPath string, logger *slog.Logger) (Competition, error) {
	files, err := os.ReadDir(folderInputPath)
	if err != nil {
		logger.Error("failed to read folder", "err", err, "folder", folderInputPath)
		return Competition{}, err
	}

	categories := []Category{}
	for _, f := range files {
		filename := fmt.Sprintf("%s/%s", folderInputPath, f.Name())
		jsonFile, err := os.Open(filename)
		if err != nil {
			logger.Error("failed to open file", "err", err, "file", filename)
			return Competition{}, err
		}
		defer jsonFile.Close()
		byteValue, err := io.ReadAll(jsonFile)
		if err != nil {
			logger.Error("failed to read file", "err", err, "file", filename)
			return Competition{}, err
		}

		var judobaseCategory Category
		err = json.Unmarshal(byteValue, &judobaseCategory)
		if err != nil {
			logger.Error("failed to unmarshal file", "err", err, "file", filename)
			return Competition{}, err
		}
		if len(judobaseCategory.Contests) > 0 && judobaseCategory.Contests[0].Weight != nil {
			judobaseCategory.Name = *judobaseCategory.Contests[0].Weight
		}

		categories = append(categories, judobaseCategory)
	}

	var competitionName string
	if len(categories) > 0 && len(categories[0].Contests) > 0 && categories[0].Contests[0].CompetitionName != nil {
		competitionName = *categories[0].Contests[0].CompetitionName
	}

	return Competition{
		Name:       competitionName,
		Categories: categories,
	}, nil
}
