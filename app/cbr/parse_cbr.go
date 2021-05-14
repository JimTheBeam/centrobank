package cbr

import (
	"centrobank/cfg"
	"log"
)

// ParseAll - parse all XMLs
func ParseAll(cfg *cfg.Config) (map[string][]Valute, error) {
	// map[date][]Valute
	result := make(map[string][]Valute, cfg.CBR.Days)

	for i := 0; i < cfg.CBR.Days; i++ {
		curr, err := parseOneXML(cfg, i)
		if err != nil {
			log.Printf("Error Failed to parse XML %v", err)
			return nil, err
		}

		result[curr.Date] = curr.Valute
	}

	return result, nil
}

// parseOneXML <n> - number of days before today
func parseOneXML(cfg *cfg.Config, n int) (ValCurs, error) {

	currency := &ValCurs{}

	// get url from config
	url := mergeUrl(cfg, getDate(n))

	// get XMLbytes
	xmlBytes, err := getXML(url, cfg)
	if err != nil {
		log.Printf("Failed to get XML: %v", err)
		return ValCurs{}, err
	}

	// decode XML from "windows-12510"
	err = unmarshalXML(xmlBytes, currency)
	if err != nil {
		return ValCurs{}, err
	}

	// add fields ValueFloat, Nominal and Cost to currency struct
	err = parseValNomCost(currency)
	if err != nil {
		return ValCurs{}, err
	}

	return *currency, nil
}
