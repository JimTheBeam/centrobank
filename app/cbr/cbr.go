package cbr

// CalculateMin - find min cost value from all valutes
func CalculateMin(data map[string][]Valute) ValuteMinMax {
	minValute := ValuteMinMax{}
	min := 10000000000000000000.0

	for i := range data {
		for _, j := range data[i] {
			if j.Cost < min {
				minValute.Date = i
				minValute.Name = j.Name
				minValute.Nominal = j.NominalInt
				minValute.Value = j.Cost

				min = j.Cost
			}
		}
	}
	return minValute
}

// CalculateMax - find max cost value from all valutes
func CalculateMax(data map[string][]Valute) ValuteMinMax {
	maxValute := ValuteMinMax{}
	max := 0.0

	for i := range data {
		for _, j := range data[i] {
			if j.Cost > max {
				maxValute.Date = i
				maxValute.Name = j.Name
				maxValute.Nominal = j.NominalInt
				maxValute.Value = j.Cost

				max = j.Cost
			}
		}
	}

	return maxValute
}
