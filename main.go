package main

import (
	"fmt"
	"github.com/andreykazakovtsev90/diploma-project/pkg/data/EmailData"
	"github.com/andreykazakovtsev90/diploma-project/pkg/data/SMSData"
	"github.com/andreykazakovtsev90/diploma-project/pkg/data/VoiceCallData"
	"github.com/andreykazakovtsev90/diploma-project/pkg/references/countryReference"
	"github.com/andreykazakovtsev90/diploma-project/pkg/references/providerReference"
	"io/ioutil"
	"log"
	"strings"
)

const countryListFilename = "./configs/countries.json"
const providerListFilename = "./configs/providers.json"
const smsDataFilename = "./simulator/sms.data"
const voiceCallDataFilename = "./simulator/voice.data"
const emailDataFilename = "./simulator/email.data"

func main() {
	// загрузка справочника стран
	if err := countryReference.Init(countryListFilename); err != nil {
		log.Fatal(err)
		return
	}

	// загрузка справочника провайдеров
	if err := providerReference.Init(providerListFilename); err != nil {
		log.Fatal(err)
		return
	}

	// Сбор данных о системе SMS
	if data, err := loadSMSData(); err != nil {
		log.Fatal(err)
		return
	} else {
		fmt.Println("Данные о системе SMS:")
		for _, d := range data {
			fmt.Println(d)
		}
	}

	// Сбор данных о системе VoiceCall
	if data, err := loadVoiceCallData(); err != nil {
		log.Fatal(err)
		return
	} else {
		fmt.Println("Данные о системе VoiceCall:")
		for _, d := range data {
			fmt.Println(d)
		}
	}

	// Сбор данных о системе Email
	if data, err := loadEmailData(); err != nil {
		log.Fatal(err)
		return
	} else {
		fmt.Println("Данные о системе Email:")
		for _, d := range data {
			fmt.Println(d)
		}
	}
}

// Сбор данных о системе SMS
func loadSMSData() ([]*SMSData.SMSData, error) {
	data := make([]*SMSData.SMSData, 0)
	file, err := ioutil.ReadFile(smsDataFilename)
	if err != nil {
		return nil, err
	}
	for _, str := range strings.Split(string(file), "\n") {
		fields := strings.Split(str, ";")
		if d, ok := SMSData.Parse(fields); ok {
			data = append(data, d)
		}
	}
	return data, nil
}

// Сбор данных о системе VoiceCall
func loadVoiceCallData() ([]*VoiceCallData.VoiceCallData, error) {
	data := make([]*VoiceCallData.VoiceCallData, 0)
	file, err := ioutil.ReadFile(voiceCallDataFilename)
	if err != nil {
		return nil, err
	}
	for _, str := range strings.Split(string(file), "\n") {
		fields := strings.Split(str, ";")
		if d, ok := VoiceCallData.Parse(fields); ok {
			data = append(data, d)
		}
	}
	return data, nil
}

// Сбор данных о системе Email
func loadEmailData() ([]*EmailData.EmailData, error) {
	data := make([]*EmailData.EmailData, 0)
	file, err := ioutil.ReadFile(emailDataFilename)
	if err != nil {
		return nil, err
	}
	for _, str := range strings.Split(string(file), "\n") {
		fields := strings.Split(str, ";")
		if d, ok := EmailData.Parse(fields); ok {
			data = append(data, d)
		}
	}
	return data, nil
}
