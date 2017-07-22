package services

import "net/url"

// Service defines a functionality provided by one or more backend servers.
type Service struct {
	Name    string
	Servers []*url.URL
}

// NewService creates a new service. It parses all provided server strings to
// URLs, returning an error if one fails to parse.
func NewService(name string, servers ...string) (*Service, error) {
	service := &Service{
		Name:    name,
		Servers: make([]*url.URL, len(servers)),
	}

	for i, server := range servers {
		u, err := url.Parse(server)
		if err != nil {
			return nil, err
		}

		service.Servers[i] = u
	}

	return service, nil
}
