package task

import (
	"testing"

	"github.com/BranDebs/Avocado-Backend/task/mocks"

	"github.com/stretchr/testify/assert"
)

func TestNewService(t *testing.T) {
	var rMock mocks.Repository

	svc := NewService(&rMock)

	assert.NotNil(t, svc, "Service cannot be nil.")
}
