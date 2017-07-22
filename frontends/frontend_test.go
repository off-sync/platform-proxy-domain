package frontends

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFrontend(t *testing.T) {
	f, err := NewFrontend("test", "test.off-sync.net", nil, "service")

	assert.NotNil(t, f)
	assert.Nil(t, err)
}

func TestNewFrontendShouldReturnErrorOnMissingFrontendName(t *testing.T) {
	f, err := NewFrontend("", "test.off-sync.net", nil, "service")

	assert.Nil(t, f)
	assert.NotNil(t, err)

	assert.Equal(t, ErrFrontendNameMissing, err)
}

func TestNewFrontendShouldReturnErrorOnMissingFrontendURL(t *testing.T) {
	f, err := NewFrontend("test", "", nil, "service")

	assert.Nil(t, f)
	assert.NotNil(t, err)

	assert.Equal(t, ErrFrontendURLMissing, err)
}

func TestNewFrontendShouldReturnErrorOnInvalidFrontendURL(t *testing.T) {
	f, err := NewFrontend("test", "://a://", nil, "service")

	assert.Nil(t, f)
	assert.NotNil(t, err)
}

func TestNewFrontendShouldReturnErrorOnMissingServiceName(t *testing.T) {
	f, err := NewFrontend("test", "test.off-sync.net", nil, "")

	assert.Nil(t, f)
	assert.NotNil(t, err)

	assert.Equal(t, ErrServiceNameMissing, err)
}
