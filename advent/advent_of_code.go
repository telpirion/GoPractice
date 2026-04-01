package misc

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func doDay1(resourceFile string) (int, int, error) {
	in, err := ioutil.ReadFile(resourceFile)
	if err != nil {
		return -1, -1, err
	}

	str := string(in)
	cals := strings.Split(str, "\n\n")
	hi := 0
	most := -1

	for i, cal := range cals {
		food := strings.Split(cal, "\n")

		tot := 0
		for _, f := range food {
			fv, err := strconv.Atoi(f)
			if err != nil {
				continue
			}
			tot += fv
		}
		if tot > hi {
			hi = tot
			most = i
		}
		tot = 0
	}
	return hi, most, nil
}

func doDay2(resourceFile string) (int, error) {
	tot := 0
	scores := make(map[string]int)
	// Could also do this as a 2-dimensional array
	scores["A X"] = 3 + 1
	scores["A Y"] = 6 + 2
	scores["A Z"] = 0 + 3
	scores["B X"] = 0 + 1
	scores["B Y"] = 3 + 2
	scores["B Z"] = 6 + 3
	scores["C X"] = 6 + 1
	scores["C Y"] = 0 + 2
	scores["C Z"] = 3 + 3

	in, err := ioutil.ReadFile("resources/day2.txt")
	if err != nil {
		return -1, err
	}

	rounds := string(in)
	rArr := strings.Split(rounds, "\n")
	for _, r := range rArr {
		rVals := strings.TrimRight(r, "\n")
		score := scores[rVals]
		tot += score
	}
	return tot, nil
}
