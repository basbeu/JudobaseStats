package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"

	"github.com/basbeu/JudobaseStats/pkg/analyser"
	"github.com/jedib0t/go-pretty/table"
)

func main() {
	inputPath := flag.String("input", "../../analysis", "path of the input folder")
	analysisFlag := flag.String("analysis", "analysis-Olympic Games Paris 2024.json,analysis-Olympic Games Tokyo 2020.json", "analysis files to compare")

	flag.Parse()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	var stats []analyser.CompetitionStats
	analysisFiles := strings.Split(*analysisFlag, ",")
	for _, fileName := range analysisFiles {
		f, err := os.Open(*inputPath + "/" + fileName)
		if err != nil {
			logger.Error("failed to open analysis file", "err", err)
			return
		}
		defer f.Close()
		byteValue, err := io.ReadAll(f)
		if err != nil {
			logger.Error("failed to read file", "err", err, "file", fileName)
			return
		}
		var s analyser.CompetitionStats
		err = json.Unmarshal(byteValue, &s)
		if err != nil {
			logger.Error("failed to unmarshal file", "err", err, "file", fileName)
			return
		}
		stats = append(stats, s)
	}

	report := map[string]analyser.RoundStats{}
	for _, competition := range stats {
		for _, category := range competition.Stats {
			if category.Category == "all" {
				for _, round := range category.Stats {
					if round.Round == "all" {
						report[competition.Name] = round
					}
				}
			}
		}
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	header := table.Row{"stat"}
	rowOrder := []string{"Fights", "Ippon", "Waza", "Yuko", "Shidos", "Hansokumake", "Unknown", "Regular", "Golden"}
	rows := map[string]table.Row{
		"Fights":      {"# Fights"},
		"Ippon":       {"# Wins by ippon"},
		"Waza":        {"# Wins by waza-ari"},
		"Yuko":        {"# Wins by yuko"},
		"Shidos":      {"# Wins by 3 shidos"},
		"Hansokumake": {"# Wins by direct hansokumake"},
		"Unknown":     {"# Unknown win type"},
		"Regular":     {"# Wins in regular time"},
		"Golden":      {"# Wins in golden score"},
	}
	for competition, stats := range report {
		header = append(header, competition)
		rows["Fights"] = append(rows["Fights"], fmt.Sprintf("%d", stats.Fights))
		rows["Ippon"] = append(rows["Ippon"], fmt.Sprintf("%d %s", stats.Ippon, formatPercentage(stats.Ippon, stats.Fights)))
		rows["Waza"] = append(rows["Waza"], fmt.Sprintf("%d %s", stats.Waza, formatPercentage(stats.Waza, stats.Fights)))
		rows["Yuko"] = append(rows["Yuko"], fmt.Sprintf("%d %s", stats.Yuko, formatPercentage(stats.Yuko, stats.Fights)))
		rows["Shidos"] = append(rows["Shidos"], fmt.Sprintf("%d %s", stats.MaxShidos, formatPercentage(stats.MaxShidos, stats.Fights)))
		rows["Hansokumake"] = append(rows["Hansokumake"], fmt.Sprintf("%d %s", stats.Hansokumake, formatPercentage(stats.Hansokumake, stats.Fights)))
		rows["Unknown"] = append(rows["Unknown"], fmt.Sprintf("%d %s", stats.Unknown, formatPercentage(stats.Unknown, stats.Fights)))
		rows["Regular"] = append(rows["Regular"], fmt.Sprintf("%d %s", stats.Regular, formatPercentage(stats.Regular, stats.Fights)))
		rows["Golden"] = append(rows["Golden"], fmt.Sprintf("%d %s", stats.Golden, formatPercentage(stats.Golden, stats.Fights)))
	}
	t.AppendHeader(header)
	for _, rowLabel := range rowOrder {
		t.AppendRow(rows[rowLabel])
	}
	t.Render()
}

func formatPercentage(part int, total int) string {
	percentage := float64(part) / float64(total) * 100
	return fmt.Sprintf("(%.2f%%)", percentage)
}
