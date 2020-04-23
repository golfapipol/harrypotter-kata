package potter_test

import (
	. "potter"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateHarryPotterBookPart1Price8Euro(t *testing.T) {
	part1Book := NewBook("Harry Potter Part 1", Euro(8.0))

	assert.Equal(t, "Harry Potter Part 1", part1Book.GetTitle())
	assert.Equal(t, Euro(8.0), part1Book.GetPrice())
}
