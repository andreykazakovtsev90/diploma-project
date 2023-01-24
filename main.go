package main

import (
	"encoding/json"
	"fmt"
	"github.com/andreykazakovtsev90/diploma-project/pkg/data/BillingData"
	"github.com/andreykazakovtsev90/diploma-project/pkg/data/EmailData"
	"github.com/andreykazakovtsev90/diploma-project/pkg/data/IncidentData"
	"github.com/andreykazakovtsev90/diploma-project/pkg/data/MMSData"
	"github.com/andreykazakovtsev90/diploma-project/pkg/data/SMSData"
	"github.com/andreykazakovtsev90/diploma-project/pkg/data/SupportData"
	"github.com/andreykazakovtsev90/diploma-project/pkg/data/VoiceCallData"
	"github.com/andreykazakovtsev90/diploma-project/pkg/references/countryReference"
	"github.com/andreykazakovtsev90/diploma-project/pkg/references/providerReference"
	"github.com/andreykazakovtsev90/diploma-project/pkg/response"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const apiURL = "127.0.0.1:8282"

var (
	countryListFilename   string
	providerListFilename  string
	smsDataFilename       string
	mmsDataURL            string
	voiceCallDataFilename string
	emailDataFilename     string
	billingDataFilename   string
	supportDataURL        string
	incidentDataURL       string
)

func ConfigInit() {
	countryListFilename = getenv("COUNTRY_LIST_FILENAME", "./configs/countries.json")
	providerListFilename = getenv("PROVIDER_LIST_FILENAME", "./configs/providers.json")
	smsDataFilename = getenv("SMS_DATA_FILENAME", "./simulator/sms.data")
	mmsDataURL = getenv("MMS_DATA_URL", "http://127.0.0.1:8383/mms")
	voiceCallDataFilename = getenv("VOICE_CALL_DATA_FILENAME", "./simulator/voice.data")
	emailDataFilename = getenv("EMAIL_DATA_FILENAME", "./simulator/email.data")
	billingDataFilename = getenv("BILLING_DATA_FILENAME", "./simulator/billing.data")
	supportDataURL = getenv("SUPPORT_DATA_FILENAME", "http://127.0.0.1:8383/support")
	incidentDataURL = getenv("INCIDENT_DATA_FILENAME", "http://127.0.0.1:8383/accendent")
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func main() {
	// инициализация переменных окружения
	ConfigInit()
	// загрузка справочника стран
	if err := countryReference.Init(countryListFilename); err != nil {
		log.Fatal(err)
	}

	// загрузка справочника провайдеров
	if err := providerReference.Init(providerListFilename); err != nil {
		log.Fatal(err)
	}

	listenAndServeHTTP()
}

func listenAndServeHTTP() {
	router := mux.NewRouter()
	router.HandleFunc("/", handleConnection)
	router.HandleFunc("/api", handleApi)
	http.ListenAndServe(apiURL, router)
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func handleApi(w http.ResponseWriter, r *http.Request) {
	resultT := new(response.ResultT)
	if resultSetT, err := getResultData(); err != nil {
		resultT.Status = false
		resultT.Error = err.Error()
		resultT.Data = nil
		res, _ := json.Marshal(resultT)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(res)
	} else {
		resultT.Status = true
		resultT.Error = ""
		resultT.Data = resultSetT
		res, _ := json.Marshal(resultT)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

// Функция сбора данных
func getResultData() (*response.ResultSetT, error) {
	resultSetT := response.NewResultSetT()

	// Сбор данных о системе SMS
	if data, err := loadSMSData(); err != nil {
		return nil, err
	} else {
		resultSetT.SetSMS(data)
	}

	// Сбор данных о системе MMS
	if data, err := loadMMSData(); err != nil {
		return nil, err
	} else {
		resultSetT.SetMMS(data)
	}

	// Сбор данных о системе VoiceCall
	if data, err := loadVoiceCallData(); err != nil {
		return nil, err
	} else {
		resultSetT.SetVoiceCall(data)
	}

	// Сбор данных о системе Email
	if data, err := loadEmailData(); err != nil {
		return nil, err
	} else {
		resultSetT.SetEmail(data)
	}

	// Сбор данных о системе Billing
	if data, err := loadBillingData(); err != nil {
		return nil, err
	} else {
		resultSetT.SetBilling(*data)
	}

	// Сбор данных о системе Support
	if data, err := loadSupportData(); err != nil {
		return nil, err
	} else {
		resultSetT.SetSupport(data)
	}

	// Сбор данных о системе истории инцидентов
	if data, err := loadIncidentData(); err != nil {
		return nil, err
	} else {
		resultSetT.SetIncidents(data)
	}
	return resultSetT, nil
}

// Сбор данных о системе SMS
func loadSMSData() ([]SMSData.SMSData, error) {
	data := make([]SMSData.SMSData, 0)
	file, err := ioutil.ReadFile(smsDataFilename)
	if err != nil {
		return nil, err
	}
	for _, str := range strings.Split(string(file), "\n") {
		fields := strings.Split(str, ";")
		if d, ok := SMSData.Parse(fields); ok {
			data = append(data, *d)
		}
	}
	return data, nil
}

// Сбор данных о системе MMS
func loadMMSData() ([]MMSData.MMSData, error) {
	data := make([]MMSData.MMSData, 0)
	res, err := http.Get(mmsDataURL)
	if err != nil {
		log.Fatal(err)
		return data, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Fatal()
		return data, fmt.Errorf("Ошибка получения данных")
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return data, err
	}
	arr := make([]MMSData.MMSData, 0)
	err = json.Unmarshal(body, &arr)
	if err != nil {
		return data, err
	}
	for _, obj := range arr {
		if obj.IsValid() {
			data = append(data, obj)
		}
	}
	return data, nil
}

// Сбор данных о системе VoiceCall
func loadVoiceCallData() ([]VoiceCallData.VoiceCallData, error) {
	data := make([]VoiceCallData.VoiceCallData, 0)
	file, err := ioutil.ReadFile(voiceCallDataFilename)
	if err != nil {
		return nil, err
	}
	for _, str := range strings.Split(string(file), "\n") {
		fields := strings.Split(str, ";")
		if d, ok := VoiceCallData.Parse(fields); ok {
			data = append(data, *d)
		}
	}
	return data, nil
}

// Сбор данных о системе Email
func loadEmailData() ([]EmailData.EmailData, error) {
	data := make([]EmailData.EmailData, 0)
	file, err := ioutil.ReadFile(emailDataFilename)
	if err != nil {
		return nil, err
	}
	for _, str := range strings.Split(string(file), "\n") {
		fields := strings.Split(str, ";")
		if d, ok := EmailData.Parse(fields); ok {
			data = append(data, *d)
		}
	}
	return data, nil
}

// Сбор данных о системе Billing
func loadBillingData() (*BillingData.BillingData, error) {
	data := BillingData.NewBillingData()
	file, err := ioutil.ReadFile(billingDataFilename)
	if err != nil {
		return nil, err
	}
	if d, err := strconv.ParseInt(string(file), 2, 8); err != nil {
		return nil, err
	} else {
		data.CreateCustomer = d&(1<<5) != 0
		data.Purchase = d&(1<<4) != 0
		data.Payout = d&(1<<3) != 0
		data.Recurring = d&(1<<2) != 0
		data.FraudControl = d&(1<<1) != 0
		data.CheckoutPage = d&(1<<0) != 0
	}
	return data, nil
}

// Сбор данных о системе Support
func loadSupportData() ([]*SupportData.SupportData, error) {
	data := make([]*SupportData.SupportData, 0)
	res, err := http.Get(supportDataURL)
	if err != nil {
		log.Fatal(err)
		return data, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Fatal()
		return data, fmt.Errorf("Ошибка получения данных")
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return data, err
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// Сбор данных о системе истории инцидентов
func loadIncidentData() ([]IncidentData.IncidentData, error) {
	data := make([]IncidentData.IncidentData, 0)
	res, err := http.Get(incidentDataURL)
	if err != nil {
		log.Fatal(err)
		return data, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Fatal()
		return data, fmt.Errorf("Ошибка получения данных")
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	arr := make([]*IncidentData.IncidentData, 0)
	err = json.Unmarshal(body, &arr)
	if err != nil {
		return nil, err
	}
	for _, obj := range arr {
		if obj.IsValid() {
			data = append(data, *obj)
		}
	}
	return data, nil
}
