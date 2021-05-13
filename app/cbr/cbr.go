package cbr

import "fmt"

// CalculateMin - find min cost value from all valutes
func CalculateMin(data map[string][]Valute) ValuteMinMax {
	minValute := ValuteMinMax{}
	min := minConst

	for date := range data {
		for _, valute := range data[date] {
			if valute.Cost < min {
				minValute.Date = date
				minValute.Name = valute.Name
				minValute.Nominal = valute.NominalInt
				minValute.Value = valute.Cost

				min = valute.Cost
			}
		}
	}
	return minValute
}

// CalculateMax - find max cost value from all valutes
func CalculateMax(data map[string][]Valute) ValuteMinMax {
	maxValute := ValuteMinMax{}
	max := maxConst

	for date := range data {
		for _, valute := range data[date] {
			if valute.Cost > max {
				maxValute.Date = date
				maxValute.Name = valute.Name
				maxValute.Nominal = valute.NominalInt
				maxValute.Value = valute.Cost

				max = valute.Cost
			}
		}
	}

	return maxValute
}

// CalculateAverage - calculate average value for all valutes
func CalculateAverage(data map[string][]Valute) {
	// map[valuteName][]values
	res := make(map[string][]float64)
	fmt.Println(len(data))
	for i := range data {
		for _, valute := range data[i] {
			res[valute.Name] = append(res[valute.Name], valute.ValueFloat)
		}
	}
	fmt.Println(len(res["SDR"]))
}

// type ValuteAverage struct {
// 	Name    string
// 	Values  []float64
// 	Average float64
// }
