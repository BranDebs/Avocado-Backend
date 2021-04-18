package task

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewService(t *testing.T) {
	var rMock MockRepository

	svc := NewService(&rMock)

	assert.NotNil(t, svc, "Service cannot be nil.")
}
