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

func ParseOneXML(cfg *cfg.Config) error {

	date := getDate(0)

	// get url from config
	url := mergeUrl(cfg, date)
	// get XMLbytes
	xmlBytes, err := GetXML(url)
	if err != nil {
		log.Printf("Failed to get XML: %v", err)
	}

	currency := &ValCurs{}

	// decode XML from "windows-12510"
	err = UnmarshalXML(xmlBytes, currency)
	if err != nil {
		return err
	}

	//TODO: Fix me Отдельная функция  ParseFloat
	ParseFloat(currency)
	fmt.Println(currency.Date)
	fmt.Println(currency.Valute)
	return nil
}

// ParseFloat - add field valueFloat to struct ValCurs
func ParseFloat(currency *ValCurs) {
	for _, el := range currency.Valute {
		value := strings.Replace(el.Value, ",", ".", 1)
		valueFloat, err := strconv.ParseFloat(value, 4)
		if err != nil {
			log.Printf("Failed to parse float %s: %v", el.Value, err)
		}
		el.ValueFloat = valueFloat
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

	log.Printf("GetXML, Status code:", resp.StatusCode)

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
