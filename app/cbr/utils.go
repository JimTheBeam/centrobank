package cbr

import (
	"bytes"
	"centrobank/cfg"
	"encoding/xml"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/html/charset"
)

// parseValNomCost - add field valueFloat, Nominal and Cost to struct ValCurs
func parseValNomCost(currency *ValCurs) error {
	for i := range currency.Valute {
		// parse ValueFloat
		valueFloat, err := strconv.ParseFloat(
			strings.Replace(currency.Valute[i].Value, ",", ".", 1),
			4)
		if err != nil {
			log.Printf("Failed to parse float %s: %v", currency.Valute[i].Value, err)
			return err
		}
		currency.Valute[i].ValueFloat = valueFloat

		// parse Nominal
		nominal, err := strconv.Atoi(currency.Valute[i].Nominal)
		if err != nil {
			log.Printf("Failed to parse Nominal %s: %v", currency.Valute[i].Nominal, err)
			return err
		}
		currency.Valute[i].NominalInt = nominal

		// parse Cost
		currency.Valute[i].Cost = currency.Valute[i].ValueFloat / float64(currency.Valute[i].NominalInt)
	}
	return nil

}

// unmarshalXML - unmarshal xmlBytes to ValCurs struct
func unmarshalXML(xmlBytes []byte, currency *ValCurs) error {
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

// mergeUrl merge host, method from config and date
// date must be like <dd/mm/yyyy>
func mergeUrl(cfg *cfg.Config, date string) string {
	return fmt.Sprintf(
		"%s/%s?date_req=%s",
		strings.TrimRight(strings.TrimLeft(cfg.CBR.Host, "/"), "/"),
		strings.TrimRight(strings.TrimLeft(cfg.CBR.Method, "/"), "/"),
		date,
	)
}

// getDate - gets date n days before now
// return string <dd/mm/yyyy> format
func getDate(n int) string {
	return time.Now().UTC().AddDate(0, 0, -n).Format(dateFormat)
}
