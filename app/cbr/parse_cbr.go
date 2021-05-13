package cbr

import (
	"bytes"
	"centrobank/cfg"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/html/charset"
)

// ParseAll - parse all XML
func ParseAll(cfg *cfg.Config) (map[string][]Valute, error) {

	result := make(map[string][]Valute)
	currency := &ValCurs{}

	for n := 0; n <= cfg.CBR.Days; n++ {
		curr, err := ParseOneXML(cfg, currency, n)
		if err != nil {
			log.Printf("Error Failed to parse XML %v", err)
			return nil, err
		}
		result[curr.Date] = curr.Valute
	}

	return result, nil
}

// ParseOneXML n - number of days before today
func ParseOneXML(cfg *cfg.Config, currency *ValCurs, n int) (*ValCurs, error) {

	date := getDate(n)

	// get url from config
	url := mergeUrl(cfg, date)

	// get XMLbytes
	xmlBytes, err := GetXML(url)
	if err != nil {
		log.Printf("Failed to get XML: %v", err)
	}

	// decode XML from "windows-12510"
	err = UnmarshalXML(xmlBytes, currency)
	if err != nil {
		return nil, err
	}

	// add float value to struct
	err = ParseFloatValue(currency)
	if err != nil {
		return nil, err
	}

	// add int nominal value to struct
	err = ParseNominal(currency)
	if err != nil {
		return nil, err
	}

	// add cost value to struct
	countCost(currency)

	return currency, nil
}

// ParseFloat - add field valueFloat to struct ValCurs
func ParseFloatValue(currency *ValCurs) error {
	for i := range currency.Valute {
		value := strings.Replace(currency.Valute[i].Value, ",", ".", 1)
		valueFloat, err := strconv.ParseFloat(value, 4)
		if err != nil {
			log.Printf("Failed to parse float %s: %v", currency.Valute[i].Value, err)
			return err
		}
		currency.Valute[i].ValueFloat = valueFloat
	}
	return nil
}

// ParseNominal - convert nominal value(string) to int
func ParseNominal(currency *ValCurs) error {
	for i := range currency.Valute {
		nominal, err := strconv.Atoi(currency.Valute[i].Nominal)
		if err != nil {
			log.Printf("Failed to parse Nominal %s: %v", currency.Valute[i].Nominal, err)
			return err
		}
		currency.Valute[i].NominalInt = nominal
	}
	return nil
}

// countCost - calculate valute cost per 1 unit
func countCost(currency *ValCurs) {
	for i := range currency.Valute {
		currency.Valute[i].Cost = currency.Valute[i].ValueFloat / float64(currency.Valute[i].NominalInt)
	}
}

// UnmarshalXML - unmarshal xmlBytes to ValCurs struct
func UnmarshalXML(xmlBytes []byte, currency *ValCurs) error {
	reader := bytes.NewReader(xmlBytes)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	err := decoder.Decode(&currency)
	if err != nil {
		log.Printf("Failed to decode XML: %v", err)
		return err
	}
	return nil
}

func GetXML(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	log.Printf("GetXML, Status code: %v", resp.StatusCode)

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("Read body: %v", err)
	}

	return data, nil
}

// mergeUrl merge host, method from config and date
// date must be like <dd/mm/yyyy>
func mergeUrl(cfg *cfg.Config, date string) string {
	return fmt.Sprintf("%s/%s?date_req=%s",
		strings.TrimRight(strings.TrimLeft(cfg.CBR.Host, "/"), "/"),
		strings.TrimRight(strings.TrimLeft(cfg.CBR.Method, "/"), "/"),
		date)
}

// getDate - gets date n days before now
// return string <dd/mm/yyyy> format
func getDate(n int) string {
	return time.Now().UTC().AddDate(0, 0, -n).Format("02/01/2006")
}
