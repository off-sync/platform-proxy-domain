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

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewService(t *testing.T) {
	s, err := NewService("name", ServiceTypeServer, "http://10.0.0.1", "http://10.0.0.2")

	assert.NotNil(t, s)
	assert.Nil(t, err)
}

func TestNewServiceShouldReturnErrorOnMissingServiceName(t *testing.T) {
	s, err := NewService("", ServiceTypeServer, "http://10.0.0.1")

	assert.Nil(t, s)
	assert.NotNil(t, err)

	assert.Equal(t, ErrServiceNameMissing, err)
}

func TestNewServiceShouldReturnErrorOnInvalidServiceType(t *testing.T) {
	s, err := NewService("name", "", "http://10.0.0.1")

	assert.Nil(t, s)
	assert.NotNil(t, err)

	assert.Equal(t, ErrInvalidServiceType, err)
}

func TestNewServiceShouldReturnErrorOnMissingServerURLs(t *testing.T) {
	s, err := NewService("name", ServiceTypeServer)

	assert.Nil(t, s)
	assert.NotNil(t, err)

	assert.Equal(t, ErrServerURLsMissing, err)
}

func TestNewServiceShouldReturnErrorOnInvalidServerURL(t *testing.T) {
	s, err := NewService("name", ServiceTypeServer, "http://10.0.0.1", "://a://")

	assert.Nil(t, s)
	assert.NotNil(t, err)
}

func TestNewServiceShouldIgnoreEmptyServerURLs(t *testing.T) {
	s, err := NewService("name", ServiceTypeServer, "", "http://10.0.0.2")

	assert.NotNil(t, s)
	assert.Nil(t, err)

	assert.Equal(t, 1, len(s.Servers))
}
