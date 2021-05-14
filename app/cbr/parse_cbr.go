package cbr

import (
	"centrobank/cfg"
	"log"
	"sync"
)

func acync(cfg *cfg.Config, i int, result map[string][]Valute, wg *sync.WaitGroup, mux *sync.Mutex) {
	defer wg.Done()

	curr, err := parseOneXML(cfg, i)
	if err != nil {
		log.Printf("Error Failed to parse XML %v", err)
	}

	mux.Lock()
	defer mux.Unlock()
	result[curr.Date] = curr.Valute

}

// ParseAll - parse all XMLs
func ParseAll(cfg *cfg.Config) (map[string][]Valute, error) {
	log.Printf("Start <ParseAll>")
	defer log.Printf("End <ParseAll>")

	// map[date][]Valute
	result := make(map[string][]Valute, cfg.CBR.Days)

	mux := &sync.Mutex{}
	wg := sync.WaitGroup{}

	for i := 0; i < cfg.CBR.Days; i++ {
		wg.Add(1)
		go acync(cfg, i, result, &wg, mux)
	}

	wg.Wait()

	return result, nil
}

// parseOneXML <n> - number of days before today
func parseOneXML(cfg *cfg.Config, n int) (ValCurs, error) {
	log.Printf("Start <parseOneXML>")
	defer log.Printf("End <parseOneXML>")

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
