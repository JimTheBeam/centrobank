package cbr

import (
	"encoding/xml"
	"math"
)

const (
	minConst   = float64(math.MaxFloat64)
	maxConst   = 0.0
	dateFormat = "02/01/2006"
)

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Date    string   `xml:"Date,attr"`
	Valute  []Valute `xml:"Valute"`
}

type Valute struct {
	ID         string `xml:"ID,attr"`
	NumCode    string `xml:"NumCode"`
	CharCode   string `xml:"CharCode"`
	Nominal    string `xml:"Nominal"`
	Name       string `xml:"Name"`
	Value      string `xml:"Value"`
	ValueFloat float64
	NominalInt int
	Cost       float64
}

type ValuteMinMax struct {
	Date    string
	Name    string
	Nominal int
	Value   float64
}

type ValuteAverage struct {
	Name    string
	Average float64
}
