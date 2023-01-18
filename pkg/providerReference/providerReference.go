package providerReference

import (
	"encoding/json"
	"io/ioutil"
)

var providerRef *ProviderReference

type ProviderReference struct {
	providers map[string]*Provider
}

type ref struct {
	Provider []*Provider `json:"providers"`
}

type Provider struct {
	Name        string `json:"name"`
	IsSMS       bool   `json:"isSMS"`
	IsMMS       bool   `json:"isMMS"`
	IsVoiceCall bool   `json:"isVoiceCall"`
	IsEmail     bool   `json:"isEmail"`
}

func Init(filename string) error {
	providerRef = new(ProviderReference)
	providerRef.providers = make(map[string]*Provider, 0)
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	r := ref{}
	err = json.Unmarshal(file, &r)
	if err != nil {
		return err
	}
	for _, provider := range r.Provider {
		providerRef.providers[provider.Name] = provider
	}
	return nil
}

func Get(name string) (*Provider, bool) {
	p, ok := providerRef.providers[name]
	return p, ok
}
