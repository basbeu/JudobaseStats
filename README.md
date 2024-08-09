# JudobaseStats

JudobaseStats is a set of tool to allow judo fans to generate their own judo statistics. 

## Prerequisites
Have a valid Go installation. It was tested with go 1.22.6.

## Scraper
The first tool is a lightweight scraper that allows to download the results of competition easily.

The scraper can be launched with the `go run` command.
```
cd JudobaseStats/cmd/scraper
go run scraper.go
```
It downloads by default the result of the `Olympic Games Paris 2024` into the `./data` folder.

To change the default values :
```
go run scraper.go -competition=IJF_COMPETITION_ID -output=FOLDER
```

### Usage
```
$ go run scraper.go -h
Usage of scraper.exe:
  -competition string
        competition ID in judobase.ijf.org (default "2653")
  -output string
        path of the output folder (default "../../data")
```

## Analyser
The second tool is an analyser. It can be used to analyse the scraped data.

As the scraper, the default values allow to analyse the `Olympic Games Paris 2024`. The flags allow to set to the desired values.

### Usage :
```
$ go run analyser.go -h
Usage of analyser.exe:
  -competition string
        competition ID in judobase.ijf.org (default "2653")
  -input string
        path of the input folder (default "../../data")
```