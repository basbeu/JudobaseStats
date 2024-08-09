package category

import (
	"errors"
	"strconv"
)

type Category int

const (
	U60             Category = 1
	U66             Category = 2
	U73             Category = 3
	U81             Category = 4
	U90             Category = 5
	U100            Category = 6
	O100            Category = 7
	U48             Category = 8
	U52             Category = 9
	U57             Category = 10
	U63             Category = 11
	U70             Category = 12
	U78             Category = 13
	O78             Category = 14
	UnknownCategory Category = 15
)

var Categories []Category
var MaleCategories []Category
var FemaleCategories []Category

func init() {
	Categories = []Category{
		U60, U66, U73, U81, U90, U100, O100, U48, U52, U57, U63, U70, U78, O78,
	}
	MaleCategories = []Category{
		U60, U66, U73, U81, U90, U100, O100,
	}
	FemaleCategories = []Category{
		U48, U52, U57, U63, U70, U78, O78,
	}
}

func (c Category) WeightID() int {
	return int(c)
}

func (c Category) String() string {
	switch c {
	case U60:
		return "-60"
	case U66:
		return "-66"
	case U73:
		return "-73"
	case U81:
		return "-81"
	case U90:
		return "-90"
	case U100:
		return "-100"
	case O100:
		return "+100"
	case U48:
		return "-48"
	case U52:
		return "-52"
	case U57:
		return "-57"
	case U63:
		return "-63"
	case U70:
		return "-70"
	case U78:
		return "-78"
	case O78:
		return "+78"
	default:
		return "unknown"
	}
}

func FromWeightID(weightID string) (Category, error) {
	id, err := strconv.ParseInt(weightID, 10, 64)
	if err != nil {
		return UnknownCategory, err
	}
	if id < int64(U60) || id >= int64(O78) {
		return UnknownCategory, errors.New("unknown category")
	}
	return Category(id), nil
}
