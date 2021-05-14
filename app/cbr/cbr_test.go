package cbr

import (
	"reflect"
	"testing"
)

func Test_calculateMin(t *testing.T) {
	type args struct {
		// data = map[date][]Valute
		data map[string][]Valute
	}
	tests := []struct {
		// name = currency name
		name string
		args args
		want ValuteMinMax
	}{
		{
			name: "USD",
			args: args{
				data: map[string][]Valute{
					"12.12.2020": {
						Valute{
							Name:       "USD",
							NominalInt: 1,
							Cost:       2.45,
						},
						Valute{
							Name:       "EURO",
							NominalInt: 1,
							Cost:       5.45,
						},
						Valute{
							Name:       "Uzbekistan Sum",
							NominalInt: 1,
							Cost:       13.45,
						},
						Valute{
							Name:       "Swiss Franc",
							NominalInt: 1,
							Cost:       16.45,
						},
						Valute{
							Name:       "China Yuan",
							NominalInt: 1,
							Cost:       28.5,
						},
					},
					"13.12.2020": {
						Valute{
							Name:       "USD",
							NominalInt: 1,
							Cost:       2.75,
						},
						Valute{
							Name:       "EURO",
							NominalInt: 1,
							Cost:       5.4,
						},
						Valute{
							Name:       "Uzbekistan Sum",
							NominalInt: 1,
							Cost:       16.45,
						},
						Valute{
							Name:       "Swiss Franc",
							NominalInt: 1,
							Cost:       14.45,
						},
						Valute{
							Name:       "China Yuan",
							NominalInt: 1,
							Cost:       29.5,
						},
					},
					"14.12.2020": {
						Valute{
							Name:       "USD",
							NominalInt: 1,
							Cost:       3.45,
						},
						Valute{
							Name:       "EURO",
							NominalInt: 1,
							Cost:       6.45,
						},
						Valute{
							Name:       "Uzbekistan Sum",
							NominalInt: 1,
							Cost:       11.45,
						},
						Valute{
							Name:       "Swiss Franc",
							NominalInt: 1,
							Cost:       19.45,
						},
						Valute{
							Name:       "China Yuan",
							NominalInt: 1,
							Cost:       32.5,
						},
					},
					"15.12.2020": {
						Valute{
							Name:       "USD",
							NominalInt: 1,
							Cost:       1.45,
						},
						Valute{
							Name:       "EURO",
							NominalInt: 1,
							Cost:       4.45,
						},
						Valute{
							Name:       "Uzbekistan Sum",
							NominalInt: 1,
							Cost:       12.45,
						},
						Valute{
							Name:       "Swiss Franc",
							NominalInt: 1,
							Cost:       15.45,
						},
						Valute{
							Name:       "China Yuan",
							NominalInt: 1,
							Cost:       27.5,
						},
					},
					"16.12.2020": {
						Valute{
							Name:       "USD",
							NominalInt: 1,
							Cost:       2.45,
						},
						Valute{
							Name:       "EURO",
							NominalInt: 1,
							Cost:       5.45,
						},
						Valute{
							Name:       "Uzbekistan Sum",
							NominalInt: 1,
							Cost:       13.45,
						},
						Valute{
							Name:       "Swiss Franc",
							NominalInt: 1,
							Cost:       16.45,
						},
						Valute{
							Name:       "China Yuan",
							NominalInt: 1,
							Cost:       28.5,
						},
					},
				},
			},

			want: ValuteMinMax{
				Date:    "15.12.2020",
				Name:    "USD",
				Nominal: 1,
				Value:   1.45,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateMin(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calculateMin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateMax(t *testing.T) {
	type args struct {
		data map[string][]Valute
	}
	tests := []struct {
		name string
		args args
		want ValuteMinMax
	}{
		{
			name: "China Yuan",
			args: args{
				data: map[string][]Valute{
					"12.12.2020": {
						Valute{
							Name:       "USD",
							NominalInt: 1,
							Cost:       2.45,
						},
						Valute{
							Name:       "EURO",
							NominalInt: 1,
							Cost:       5.45,
						},
						Valute{
							Name:       "Uzbekistan Sum",
							NominalInt: 1,
							Cost:       13.45,
						},
						Valute{
							Name:       "Swiss Franc",
							NominalInt: 1,
							Cost:       16.45,
						},
						Valute{
							Name:       "China Yuan",
							NominalInt: 1,
							Cost:       28.5,
						},
					},
					"13.12.2020": {
						Valute{
							Name:       "USD",
							NominalInt: 1,
							Cost:       2.75,
						},
						Valute{
							Name:       "EURO",
							NominalInt: 1,
							Cost:       5.4,
						},
						Valute{
							Name:       "Uzbekistan Sum",
							NominalInt: 1,
							Cost:       16.45,
						},
						Valute{
							Name:       "Swiss Franc",
							NominalInt: 1,
							Cost:       14.45,
						},
						Valute{
							Name:       "China Yuan",
							NominalInt: 1,
							Cost:       29.5,
						},
					},
					"14.12.2020": {
						Valute{
							Name:       "USD",
							NominalInt: 1,
							Cost:       3.45,
						},
						Valute{
							Name:       "EURO",
							NominalInt: 1,
							Cost:       6.45,
						},
						Valute{
							Name:       "Uzbekistan Sum",
							NominalInt: 1,
							Cost:       11.45,
						},
						Valute{
							Name:       "Swiss Franc",
							NominalInt: 1,
							Cost:       19.45,
						},
						Valute{
							Name:       "China Yuan",
							NominalInt: 1,
							Cost:       32.5,
						},
					},
					"15.12.2020": {
						Valute{
							Name:       "USD",
							NominalInt: 1,
							Cost:       1.45,
						},
						Valute{
							Name:       "EURO",
							NominalInt: 1,
							Cost:       4.45,
						},
						Valute{
							Name:       "Uzbekistan Sum",
							NominalInt: 1,
							Cost:       12.45,
						},
						Valute{
							Name:       "Swiss Franc",
							NominalInt: 1,
							Cost:       15.45,
						},
						Valute{
							Name:       "China Yuan",
							NominalInt: 1,
							Cost:       27.5,
						},
					},
					"16.12.2020": {
						Valute{
							Name:       "USD",
							NominalInt: 1,
							Cost:       2.45,
						},
						Valute{
							Name:       "EURO",
							NominalInt: 1,
							Cost:       5.45,
						},
						Valute{
							Name:       "Uzbekistan Sum",
							NominalInt: 1,
							Cost:       13.45,
						},
						Valute{
							Name:       "Swiss Franc",
							NominalInt: 1,
							Cost:       16.45,
						},
						Valute{
							Name:       "China Yuan",
							NominalInt: 1,
							Cost:       28.5,
						},
					},
				},
			},

			want: ValuteMinMax{
				Date:    "14.12.2020",
				Name:    "China Yuan",
				Nominal: 1,
				Value:   32.5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateMax(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calculateMax() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateAverage(t *testing.T) {
	type args struct {
		data map[string][]Valute
	}
	tests := []struct {
		name string
		args args
		want []ValuteAverage
	}{
		{
			name: "Average",
			args: args{
				data: map[string][]Valute{
					"12.12.2020": {
						Valute{
							Name:       "USD",
							NominalInt: 1,
							Cost:       2.45,
						},
						Valute{
							Name:       "EURO",
							NominalInt: 1,
							Cost:       5.45,
						},
						Valute{
							Name:       "Uzbekistan Sum",
							NominalInt: 1,
							Cost:       13.45,
						},
						Valute{
							Name:       "Swiss Franc",
							NominalInt: 1,
							Cost:       16.45,
						},
						Valute{
							Name:       "China Yuan",
							NominalInt: 1,
							Cost:       28.5,
						},
					},

					"13.12.2020": {
						Valute{
							Name:       "USD",
							NominalInt: 1,
							Cost:       2.75,
						},
						Valute{
							Name:       "EURO",
							NominalInt: 1,
							Cost:       5.4,
						},
						Valute{
							Name:       "Uzbekistan Sum",
							NominalInt: 1,
							Cost:       16.45,
						},
						Valute{
							Name:       "Swiss Franc",
							NominalInt: 1,
							Cost:       14.45,
						},
						Valute{
							Name:       "China Yuan",
							NominalInt: 1,
							Cost:       29.5,
						},
					},

					"14.12.2020": {
						Valute{
							Name:       "USD",
							NominalInt: 1,
							Cost:       3.45,
						},
						Valute{
							Name:       "EURO",
							NominalInt: 1,
							Cost:       6.45,
						},
						Valute{
							Name:       "Uzbekistan Sum",
							NominalInt: 1,
							Cost:       11.45,
						},
						Valute{
							Name:       "Swiss Franc",
							NominalInt: 1,
							Cost:       19.45,
						},
						Valute{
							Name:       "China Yuan",
							NominalInt: 1,
							Cost:       32.5,
						},
					},

					"15.12.2020": {
						Valute{
							Name:       "USD",
							NominalInt: 1,
							Cost:       1.45,
						},
						Valute{
							Name:       "EURO",
							NominalInt: 1,
							Cost:       4.45,
						},
						Valute{
							Name:       "Uzbekistan Sum",
							NominalInt: 1,
							Cost:       12.45,
						},
						Valute{
							Name:       "Swiss Franc",
							NominalInt: 1,
							Cost:       15.45,
						},
						Valute{
							Name:       "China Yuan",
							NominalInt: 1,
							Cost:       27.5,
						},
					},

					"16.12.2020": {
						Valute{
							Name:       "USD",
							NominalInt: 1,
							Cost:       2.45,
						},
						Valute{
							Name:       "EURO",
							NominalInt: 1,
							Cost:       5.45,
						},
						Valute{
							Name:       "Uzbekistan Sum",
							NominalInt: 1,
							Cost:       13.45,
						},
						Valute{
							Name:       "Swiss Franc",
							NominalInt: 1,
							Cost:       16.45,
						},
						Valute{
							Name:       "China Yuan",
							NominalInt: 1,
							Cost:       28.5,
						},
					},
				},
			},

			want: []ValuteAverage{
				{
					Name:    "USD",
					Average: 2.51,
				},
				{
					Name:    "EURO",
					Average: 5.44,
				},
				{
					Name:    "Uzbekistan Sum",
					Average: 13.45,
				},
				{
					Name:    "Swiss Franc",
					Average: 16.45,
				},
				{
					Name:    "China Yuan",
					Average: 29.3,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := calculateAverage(tt.args.data); !reflect.DeepEqual(selectionSort(got), tt.want) {
				t.Errorf("calculateAverage() = %v, want %v", got, tt.want)
			}
		})
	}
}

// selectionSort - sort []ValuteAverage by ValuteAverage value
func selectionSort(arr []ValuteAverage) []ValuteAverage {
	var low int // index of a lowest number

	for n := 0; n < len(arr); n++ {

		low = n
		for i := n + 1; i < len(arr); i++ {
			if arr[low].Average > arr[i].Average {
				low = i
			}
		}
		arr[n], arr[low] = arr[low], arr[n]
	}
	return arr
}
