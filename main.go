package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/basbeu/JudobaseStats/judobase"
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
	fmt.Println(judobaseComp)
}
