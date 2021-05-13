package cbr

import (
	"encoding/xml"
)

// TODO: probably delete
type Currencies struct {
	Date     string
	DateData []ValCurs
}

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
}
