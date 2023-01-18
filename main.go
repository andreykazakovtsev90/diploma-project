package main

import (
	"fmt"
	"github.com/andreykazakovtsev90/diploma-project/pkg/SMSData"
	"github.com/andreykazakovtsev90/diploma-project/pkg/countryRef"
	"io/ioutil"
	"log"
	"strings"
)

const countryListFilename = "./configs/countries.json"
const smsDataFilename = "./simulator/sms.data"

var providers = map[string]bool{"Topolo": true, "Rond": true, "Kildy": true}

var ref *countryRef.CountryRef

// Сбор данных о системе SMS
func loadSMSData() ([]*SMSData.SMSData, error) {
	data := make([]*SMSData.SMSData, 0)
	file, err := ioutil.ReadFile(smsDataFilename)
	if err != nil {
		return nil, err
	}
	for _, str := range strings.Split(string(file), "\n") {
		fields := strings.Split(str, ";")
		if !validateSMSData(fields) {
			continue
		}
		d := SMSData.NewSMSData(fields[0], fields[1], fields[2], fields[3])
		data = append(data, d)
	}
	return data, nil
}

func validateSMSData(fields []string) bool {
	if len(fields) != 4 {
		return false
	}
	if !ref.Contains(fields[0]) {
		return false
	}
	if !providers[fields[3]] {
		return false
	}
	return true
}

func main() {
	// загрузка справочника стран
	ref = countryRef.NewCountryRef()
	err := ref.Init(countryListFilename)
	if err != nil {
		log.Fatal(err)
		return
	}
	if data, err := loadSMSData(); err != nil {
		log.Fatal(err)
		return
	} else {
		fmt.Println("Данных о системе SMS:")
		for _, d := range data {
			fmt.Println(d)
		}
	}
}
