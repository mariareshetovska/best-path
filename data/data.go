package data

import (
	"best-path/utility"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TimeTable struct {
	Name           string
	LeaveSt        int
	ArriveSt       int
	Price          float64
	LeaveTime      string
	ArriveTime     string
	TravelDuration int
}

func ReadFile(path string) ([][]string, error) {
	csvFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	return csvLines, err

}

func MapData(csvLines [][]string) []TimeTable {
	var data []TimeTable
	for i := range csvLines {
		for j := range csvLines[i] {
			s := strings.Split(csvLines[i][j], ";")
			table := TimeTable{
				Name:       s[0],
				LeaveTime:  s[4],
				ArriveTime: s[5],
			}
			leaveSt, err := strconv.Atoi(s[1])
			if err != nil {
				fmt.Println(err)
			}
			table.LeaveSt = leaveSt
			arriveSt, err := strconv.Atoi(s[2])
			utility.CheckError(err)

			table.ArriveSt = arriveSt
			price, _ := strconv.ParseFloat(s[3], 64)
			table.Price = price

			t1, err := utility.ParseDurationToSec(s[4])
			utility.CheckError(err)

			t2, err := utility.ParseDurationToSec(s[5])
			utility.CheckError(err)

			table.TravelDuration = utility.GetDuration(t1, t2)
			data = append(data, table)
		}
	}
	return data
}
