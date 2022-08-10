package utility

import "fmt"

func FindMin(a []int) (min int) {
	min = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
	}
	return min
}

func FindMinFl(a []float64) (min float64) {
	min = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
	}
	return min
}

func Unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func ParseDurationToSec(st string) (int, error) {
	var h, m, s int
	n, err := fmt.Sscanf(st, "%d:%d:%d", &h, &m, &s)
	if err != nil || n != 3 {
		return 0, err
	}
	return h*3600 + m*60 + s, nil
}

func ParseDurationToHour(totalSecs int) string {
	h := totalSecs / 3600
	m := (totalSecs % 3600) / 60
	s := totalSecs % 60
	return fmt.Sprintf("%dh %dm %ds", h, m, s)
}

func GetDuration(t1, t2 int) int {
	if t1 > t2 {
		return 86400 - t1 + t2
	}
	return t2 - t1
}

func CheckError(err error) error {
	if err != nil {
		return err
	}
	return nil
}

func FillInfInt(matrix [][]int, inf int) [][]int {
	for i := range matrix {
		for j := range matrix {
			if matrix[i][j] == 0 {
				matrix[i][j] = inf
			}
			if i == j {
				matrix[i][j] = 0
			}
		}
	}
	return matrix

}
func FillInfFl(matrix [][]float64, inf float64) [][]float64 {
	for i := range matrix {
		for j := range matrix {
			if matrix[i][j] == 0 {
				matrix[i][j] = inf
			}
			if i == j {
				matrix[i][j] = 0
			}
		}
	}
	return matrix

}
func FillInfFloat(matrix [][]float64, inf float64) [][]float64 {
	for i := range matrix {
		for j := range matrix {
			if matrix[i][j] == 0 {
				matrix[i][j] = inf
			}
			if i == j {
				matrix[i][j] = 0
			}
		}
	}
	return matrix

}
