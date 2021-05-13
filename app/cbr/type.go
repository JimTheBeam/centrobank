package cbr

import (
	"encoding/xml"
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
