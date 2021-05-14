package cbr

import (
	"centrobank/cfg"
	"fmt"
	"log"
)

func PrintResults(data map[string][]Valute, cfg *cfg.Config) {
	log.Printf("Start <PrintResults>")
	defer log.Printf("End <PrintResults>")

	// calculate min value currency
	min := calculateMin(data)
	fmt.Printf(
		"The currency with the MINIMUM value is %s\n",
		min.Name,
	)
	fmt.Printf(
		"The minimum value was reached on %s. One %s cost %.6f rubles.\n\n",
		min.Date,
		min.Name,
		min.Value,
	)

	// calculate max value currency
	max := calculateMax(data)
	fmt.Printf(
		"The currency with the MAXIMUM value is %s\n",
		max.Name,
	)
	fmt.Printf(
		"The maximum value was reached on %s. One %s cost %.4f rubles.\n\n",
		max.Date,
		max.Name,
		max.Value,
	)

	// calculate average value of all currencies
	res := calculateAverage(data)
	fmt.Printf("Average currency values for the last %d days:\n", cfg.CBR.Days)
	for _, valutes := range res {
		fmt.Printf("%s:  %.4f rubles.\n", valutes.Name, valutes.Average)
	}
}

// calculateMin - find min cost value from all valutes
func calculateMin(data map[string][]Valute) ValuteMinMax {
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

// calculateMax - find max cost value from all valutes
func calculateMax(data map[string][]Valute) ValuteMinMax {
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

// calculateAverage - calculate average value for all valutes
func calculateAverage(data map[string][]Valute) []ValuteAverage {
	// map[valuteName][]currencyValues
	res := make(map[string][]float64, len(data))

	for date := range data {
		for _, valute := range data[date] {
			res[valute.Name] = append(res[valute.Name], valute.ValueFloat)
		}
	}

	averageValutes := make([]ValuteAverage, 0, len(res))

	for name, values := range res {
		sum := 0.0
		for i := range values {
			sum += values[i]
		}

		valAv := ValuteAverage{
			Name:    name,
			Average: sum / float64(len(values)),
		}

		averageValutes = append(averageValutes, valAv)
	}
	return averageValutes
}
