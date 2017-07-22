package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewService(t *testing.T) {
	s, err := NewService("name", "http://10.0.0.1", "http://10.0.0.2")

	assert.NotNil(t, s)
	assert.Nil(t, err)
}

func TestNewServiceShouldReturnErrorOnMissingServiceName(t *testing.T) {
	s, err := NewService("", "http://10.0.0.1")

	assert.Nil(t, s)
	assert.NotNil(t, err)

	assert.Equal(t, ErrServiceNameMissing, err)
}

func TestNewServiceShouldReturnErrorOnMissingServerURLs(t *testing.T) {
	s, err := NewService("name")

	assert.Nil(t, s)
	assert.NotNil(t, err)

	assert.Equal(t, ErrServerURLsMissing, err)
}

func TestNewServiceShouldReturnErrorOnInvalidServerURL(t *testing.T) {
	s, err := NewService("name", "http://10.0.0.1", "://a://")

	assert.Nil(t, s)
	assert.NotNil(t, err)
}
