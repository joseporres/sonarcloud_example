package example_test

import (
	"testing"

	"github.com/joseporres/sonarcloud_example/example"
	"github.com/stretchr/testify/assert"
)

func TestCreateMessage(t *testing.T) {
	m, err := example.CreateMessage("Shoichi")
	assert.Nil(t, err)
	assert.Equal(t, "Hello, Shoichi", m)
}
