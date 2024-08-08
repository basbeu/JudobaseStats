package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/basbeu/JudobaseStats/judobase"
	"github.com/basbeu/JudobaseStats/report"
	"github.com/basbeu/JudobaseStats/stats"
)

func main() {
	input := flag.String("input", "./input.json", "input json file from judobase.ijf.org")

	flag.Parse()

	jsonFile, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
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

	winRecords := stats.ParseWinRecords(judobaseComp)
	winByTypes := report.GroupByWinType(winRecords)
	winByGolden := report.GroupByGoldenScore(winRecords)

	fmt.Println("====================================")
	fmt.Printf("# of fights: %d\n", len(winRecords))
	fmt.Println("====================================")
	fmt.Printf("Wins by ippon: %d %s\n", winByTypes[stats.WinByIppon], formatPercentage(winByTypes[stats.WinByIppon], len(winRecords)))
	fmt.Printf("Wins by waza-ari: %d %s\n", winByTypes[stats.WinByWaza], formatPercentage(winByTypes[stats.WinByWaza], len(winRecords)))
	fmt.Printf("Wins by 3 shidos: %d %s\n", winByTypes[stats.WinByShido], formatPercentage(winByTypes[stats.WinByShido], len(winRecords)))
	fmt.Printf("Wins by direct hansoku-make: %d %s\n", winByTypes[stats.WinByHansokuMake], formatPercentage(winByTypes[stats.WinByHansokuMake], len(winRecords)))
	fmt.Println("====================================")
	fmt.Printf("Wins in regular time: %d %s\n", winByGolden[false], formatPercentage(winByGolden[false], len(winRecords)))
	fmt.Printf("Wins in Golden Score: %d %s\n", winByGolden[true], formatPercentage(winByGolden[true], len(winRecords)))
	fmt.Println("====================================")
}

func formatPercentage(part int, total int) string {
	percentage := float64(part) / float64(total) * 100
	return fmt.Sprintf("(%.2f%%)", percentage)
}
