package config

import "os"

type ServiceList struct {
	Service []string
}

func NewServiceList() *ServiceList {
	return &ServiceList{
		Service: []string{"log"},
	}
}

func (s *ServiceList) GetServiceURL(name string) string {
	for _, service := range s.Service {
		if service == name {
			return os.Getenv(service)
		}
	}
	return ""
}
