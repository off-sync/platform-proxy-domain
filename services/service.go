// Copyright (c) 2017 off-sync
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

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
func NewService(name string, serviceType ServiceType, serverURLs ...string) (*Service, error) {
	if name == "" {
		return nil, ErrServiceNameMissing
	}

	if serviceType != ServiceTypeServer &&
		serviceType != ServiceTypeSystem &&
		serviceType != ServiceTypeUnsupported {
		return nil, ErrInvalidServiceType
	}

	if len(serverURLs) < 1 {
		return nil, ErrServerURLsMissing
	}

	service := &Service{
		Name: name,
		Type: serviceType,
	}

	for _, serverURL := range serverURLs {
		if serverURL == "" {
			continue
		}

		u, err := url.Parse(serverURL)
		if err != nil {
			return nil, err
		}

		service.Servers = append(service.Servers, u)
	}

	return service, nil
}
