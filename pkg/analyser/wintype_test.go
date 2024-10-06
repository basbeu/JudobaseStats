package analyser

import (
	"testing"

	"github.com/basbeu/JudobaseStats/internal/judobase"
	"github.com/stretchr/testify/assert"
)

func generateWinTypes(isWinnerWhite bool) []map[winType]judobase.Contest {
	winTypes := []map[winType]judobase.Contest{
		{
			winByIppon: {
				IpponWhite: whiteScore(isWinnerWhite, "1", "0"),
				IpponBlue:  blueScore(isWinnerWhite, "1", "0"),
			},
		},
		{
			winByIppon: {
				IpponWhite: whiteScore(isWinnerWhite, "1", "0"),
				WazaWhite:  whiteScore(isWinnerWhite, "1", "0"),
				IpponBlue:  blueScore(isWinnerWhite, "1", "0"),
				WazaBlue:   blueScore(isWinnerWhite, "1", "0"),
			},
		},
		{
			winByIppon: {
				IpponWhite: whiteScore(isWinnerWhite, "1", "1"),
				WazaWhite:  whiteScore(isWinnerWhite, "1", "1"),
				IpponBlue:  blueScore(isWinnerWhite, "1", "1"),
				WazaBlue:   blueScore(isWinnerWhite, "1", "1"),
			},
		},
		{
			winByIppon: {
				WazaWhite: whiteScore(isWinnerWhite, "2", "0"),
				WazaBlue:  blueScore(isWinnerWhite, "2", "0"),
			},
		},
		{
			winByIppon: {
				WazaWhite: whiteScore(isWinnerWhite, "2", "1"),
				WazaBlue:  blueScore(isWinnerWhite, "2", "1"),
			},
		},
		{
			winByWaza: {
				WazaWhite: whiteScore(isWinnerWhite, "1", "0"),
				WazaBlue:  blueScore(isWinnerWhite, "1", "0"),
			},
		},
		{
			winBy3Shidos: {
				IpponWhite:   whiteScore(isWinnerWhite, "1", "0"),
				IpponBlue:    blueScore(isWinnerWhite, "1", "0"),
				PenaltyWhite: whiteScore(isWinnerWhite, "0", "3"),
				PenaltyBlue:  blueScore(isWinnerWhite, "0", "3"),
			},
		},
		{
			winBy3Shidos: {
				IpponWhite:   whiteScore(isWinnerWhite, "1", "0"),
				IpponBlue:    blueScore(isWinnerWhite, "1", "0"),
				PenaltyWhite: whiteScore(isWinnerWhite, "1", "3"),
				PenaltyBlue:  blueScore(isWinnerWhite, "1", "3"),
			},
		},
		{
			winBy3Shidos: {
				IpponWhite:   whiteScore(isWinnerWhite, "1", "0"),
				IpponBlue:    blueScore(isWinnerWhite, "1", "0"),
				PenaltyWhite: whiteScore(isWinnerWhite, "2", "3"),
				PenaltyBlue:  blueScore(isWinnerWhite, "2", "3"),
			},
		},
		{
			winByHansokuMake: {
				IpponWhite: whiteScore(isWinnerWhite, "1", "0"),
				IpponBlue:  blueScore(isWinnerWhite, "1", "0"),
				HSKWhite:   whiteScore(isWinnerWhite, "0", "1"),
				HSKBlue:    blueScore(isWinnerWhite, "0", "1"),
			},
		},
		{
			winByHansokuMake: {
				IpponWhite: whiteScore(isWinnerWhite, "1", "1"),
				WazaWhite:  whiteScore(isWinnerWhite, "1", "1"),
				IpponBlue:  blueScore(isWinnerWhite, "1", "1"),
				WazaBlue:   blueScore(isWinnerWhite, "1", "1"),
				HSKWhite:   whiteScore(isWinnerWhite, "0", "1"),
				HSKBlue:    blueScore(isWinnerWhite, "0", "1"),
			},
		},
		{
			winByHansokuMake: {
				IpponWhite:   whiteScore(isWinnerWhite, "1", "1"),
				WazaWhite:    whiteScore(isWinnerWhite, "1", "1"),
				IpponBlue:    blueScore(isWinnerWhite, "1", "1"),
				WazaBlue:     blueScore(isWinnerWhite, "1", "1"),
				HSKWhite:     whiteScore(isWinnerWhite, "0", "1"),
				HSKBlue:      blueScore(isWinnerWhite, "0", "1"),
				PenaltyWhite: whiteScore(isWinnerWhite, "0", "2"),
				PenaltyBlue:  blueScore(isWinnerWhite, "0", "2"),
			},
		},
	}

	IDWinner := strPtr("blue")
	if isWinnerWhite {
		IDWinner = strPtr("white")
	}

	for i, winType := range winTypes {
		for expected, contest := range winType {
			contest.IDWinner = IDWinner
			contest.IDPersonWhite = strPtr("white")
			contest.IDPersonBlue = strPtr("blue")
			winTypes[i][expected] = contest
		}
	}

	return winTypes
}

func whiteScore(isWinnerWhite bool, winningValue string, losingValue string) *string {
	if isWinnerWhite {
		return &winningValue
	}
	return &losingValue
}

func blueScore(isWinnerWhite bool, winningValue string, losingValue string) *string {
	if !isWinnerWhite {
		return &winningValue
	}
	return &losingValue
}

func TestParseWinType(t *testing.T) {
	cases := generateWinTypes(true)
	cases = append(cases, generateWinTypes(false)...)

	for _, testCase := range cases {
		for expected, contest := range testCase {
			assert.Equal(t, expected, parseWinType(contest))
		}
	}
}

func strPtr(s string) *string {
	return &s
}
