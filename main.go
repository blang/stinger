package main

import (
	"flag"
	"log"
	"net/http"
)

type Overlay struct {
	Name      string      `json:"name"`
	Instances []*Instance `json:"instances"`
}

type Instance struct {
	Name      string `json:"name"`
	Game      string `json:"game"`
	Modstring string `json:"modstring"`
	Betamod   string `json:"betamod"`
	Host      string `json:"host"`
	Port      string `json:"port"`
	Password  string `json:"password"`
}

type Provider interface {
	Overlays() []*Overlay
}

func main() {
	var (
		listen     = flag.String("listen", ":9000", "Public interface, e.g. 127.0.0.1:9000")
		configFile = flag.String("config", "config.json", "Config file")
	)
	flag.Parse()

	conf, err := ReadConfigFromFile(*configFile)
	if err != nil {
		log.Fatalf("Could not read config from file %q, error: %q", *configFile, err)
		return
	}

	provider := ProviderProxyFromConfig(conf)
	restapi := NewRestAPI(provider)
	log.Printf("Start RestAPI listening on %q", *listen)
	if err := http.ListenAndServe(*listen, restapi); err != nil {
		log.Fatalf("HTTP Server crashed: %q", err)
		return
	}
}

func ProviderProxyFromConfig(conf *Config) Provider {
	services := make([]Provider, 0, len(conf.MantisServices))
	for _, conf := range conf.MantisServices {
		service := NewMantisService(conf.Host, conf.Key)
		services = append(services, service)
	}
	return NewProviderProxy(services)
}
