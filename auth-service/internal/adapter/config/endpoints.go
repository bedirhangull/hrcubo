package config

import (
	"errors"
	"os"
)

type ServiceList struct {
	services map[string]string
}

func NewServiceList() *ServiceList {
	return &ServiceList{
		services: map[string]string{
			"log": "localhost:8081",
		},
	}
}

func (s *ServiceList) ResolveServiceURL(name string) (string, error) {
	url, ok := s.services[name]
	if !ok {
		return "", errors.New("service not found: " + name)
	}

	if os.Getenv("APP_ENV") == "production" {
		prodURL := os.Getenv(name + "_SERVICE_URL")
		if prodURL != "" {
			return prodURL, nil
		}
	}

	return url, nil
}
