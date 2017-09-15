package services

import "net/url"

// ServiceType is used to described the types of services.
type ServiceType string

// Possible ServiceType values.
const (
	ServiceTypeServer      ServiceType = "server"
	ServiceTypeSystem      ServiceType = "system"
	ServiceTypeUnsupported ServiceType = "unsupported"
)

// Service defines a functionality provided by one or more backend servers.
type Service struct {
	Name    string
	Type    ServiceType
	Servers []*url.URL
}

// NewService creates a new service. It parses all provided server strings to
// URLs, returning an error if one fails to parse.
func NewService(name string, serverURLs ...string) (*Service, error) {
	if name == "" {
		return nil, ErrServiceNameMissing
	}

	if len(serverURLs) < 1 {
		return nil, ErrServerURLsMissing
	}

	service := &Service{
		Name:    name,
		Servers: make([]*url.URL, len(serverURLs)),
	}

	for i, serverURL := range serverURLs {
		u, err := url.Parse(serverURL)
		if err != nil {
			return nil, err
		}

		service.Servers[i] = u
	}

	return service, nil
}
