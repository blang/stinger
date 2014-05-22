package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type ProviderProxy struct {
	Services []Provider
}

func NewProviderProxy(services []Provider) *ProviderProxy {
	return &ProviderProxy{
		Services: services,
	}
}

func (p *ProviderProxy) Overlays() []*Overlay {
	var overlays []*Overlay
	for _, service := range p.Services {
		sOverlays := service.Overlays()
		if sOverlays != nil {
			overlays = append(overlays, sOverlays...)
		}
	}
	return overlays
}

type MantisPrivInstanceList struct {
	Instances    []*MantisPrivInstance `json:"instances"`
	ExtInstances []*MantisPrivInstance `json:"extinstances"`
}

type MantisPrivInstance struct {
	Name      string `json:"name"`
	Game      string `json:"game"`
	Modstring string `json:"modstring"`
	Betamod   string `json:"betamod"`
	Host      string `json:"host"`
	Port      int    `json:"port"`
	Password  string `json:"password"`
}

type MantisService struct {
	Host string
	Key  string
}

func NewMantisService(host string, key string) *MantisService {
	return &MantisService{
		Host: host,
		Key:  key,
	}
}

func (s *MantisService) Overlays() []*Overlay {
	req, err := http.NewRequest("GET", "http://"+s.Host+"/instances", nil)
	if err != nil {
		log.Printf("Could not create request")
		return nil
	}
	req.Header.Set("API-KEY", s.Key)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Could not reach host, error: %q", err)
		return nil
	}
	if resp.StatusCode != http.StatusOK {
		log.Printf("Statuscode was %d", resp.StatusCode)
		return nil
	}
	var privInstanceList MantisPrivInstanceList
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Could not read response body")
		return nil
	}
	err = json.Unmarshal(b, &privInstanceList)
	if err != nil {
		log.Printf("Could not unmarshal response from %s", s.Host)
		return nil
	}

	instOverlay := &Overlay{
		Name: "Internal Instances from " + s.Host,
	}
	for _, inst := range privInstanceList.Instances {
		overlayInst := &Instance{
			Name:      inst.Name,
			Game:      inst.Game,
			Modstring: inst.Modstring,
			Betamod:   inst.Betamod,
			Host:      inst.Host,
			Port:      inst.Port,
			Password:  inst.Password,
		}
		instOverlay.Instances = append(instOverlay.Instances, overlayInst)
	}

	extInstOverlay := &Overlay{
		Name: "External Instances from " + s.Host,
	}
	for _, inst := range privInstanceList.ExtInstances {
		overlayInst := &Instance{
			Name:      inst.Name,
			Game:      inst.Game,
			Modstring: inst.Modstring,
			Betamod:   inst.Betamod,
			Host:      inst.Host,
			Port:      inst.Port,
			Password:  inst.Password,
		}
		extInstOverlay.Instances = append(extInstOverlay.Instances, overlayInst)
	}
	return []*Overlay{
		instOverlay,
		extInstOverlay,
	}
}
