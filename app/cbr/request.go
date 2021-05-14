package cbr

import (
	"centrobank/cfg"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// getXML - gets XML from cbr
func getXML(url string, cfg *cfg.Config) ([]byte, error) {

	client := &http.Client{Timeout: cfg.CBR.Timeout * time.Second}

	resp, err := client.Get(url)
	if err != nil {
		return []byte{}, fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	log.Printf("GetXML, Status code: %v, Url: %s", resp.StatusCode, url)

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("Read body: %v", err)
	}

	return data, nil
}
