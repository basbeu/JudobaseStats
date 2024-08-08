package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/basbeu/JudobaseStats/internal/judobase"
	"github.com/basbeu/JudobaseStats/pkg/analyser"
)

func main() {
	input := flag.String("input", "./input.json", "input json file from judobase.ijf.org")

	flag.Parse()

	jsonFile, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
		return
	}

	var judobaseComp judobase.Competition
	err = json.Unmarshal(byteValue, &judobaseComp)
	if err != nil {
		log.Fatal(err)
		return
	}

	winRecords := analyser.ParseWinRecords(judobaseComp)
	winByTypes := analyser.GroupByWinType(winRecords)
	winByGolden := analyser.GroupByGoldenScore(winRecords)

	fmt.Println("====================================")
	fmt.Printf("# of fights: %d\n", len(winRecords))
	fmt.Println("====================================")
	fmt.Printf("Wins by ippon: %d %s\n", winByTypes[analyser.WinByIppon], formatPercentage(winByTypes[analyser.WinByIppon], len(winRecords)))
	fmt.Printf("Wins by waza-ari: %d %s\n", winByTypes[analyser.WinByWaza], formatPercentage(winByTypes[analyser.WinByWaza], len(winRecords)))
	fmt.Printf("Wins by 3 shidos: %d %s\n", winByTypes[analyser.WinByShido], formatPercentage(winByTypes[analyser.WinByShido], len(winRecords)))
	fmt.Printf("Wins by direct hansoku-make: %d %s\n", winByTypes[analyser.WinByHansokuMake], formatPercentage(winByTypes[analyser.WinByHansokuMake], len(winRecords)))
	fmt.Println("====================================")
	fmt.Printf("Wins in regular time: %d %s\n", winByGolden[false], formatPercentage(winByGolden[false], len(winRecords)))
	fmt.Printf("Wins in Golden Score: %d %s\n", winByGolden[true], formatPercentage(winByGolden[true], len(winRecords)))
	fmt.Println("====================================")
}

func formatPercentage(part int, total int) string {
	percentage := float64(part) / float64(total) * 100
	return fmt.Sprintf("(%.2f%%)", percentage)
}
