package main

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
		overlays = append(overlays, service.Overlays()...)
	}
	return overlays
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
	return []*Overlay{
		&Overlay{
			Name: "My Overlay for Host" + s.Host,
			Instances: []*Instance{
				&Instance{
					Name: "Instance Name",
				},
			},
		},
	}
}
