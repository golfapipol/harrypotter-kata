package potter_test

import (
	. "potter"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectPotterPart1BookIntoBasket(t *testing.T) {
	part1Book := NewBook("Harry Potter Part 1", Euro(8.00))
	customer := NewCustomer()

	customer.SelectToBasket(part1Book, 1)

	currentBasket := customer.GetBasket()
	assert.Equal(t, 1, currentBasket.CountBook())
	assert.Equal(t, Euro(8.00), currentBasket.GetTotalPrice())
	assert.Equal(t, Euro(0.00), currentBasket.GetDiscountPrice())
	assert.Equal(t, Euro(8.00), currentBasket.GetNetPrice())

}

func TestSelectPotterPart1BookAndPotterPart2BookIntoBasket(t *testing.T) {
	part1Book := NewBook("Harry Potter Part 1", Euro(8.00))
	part2Book := NewBook("Harry Potter Part 2", Euro(8.00))
	customer := NewCustomer()

	customer.SelectToBasket(part1Book, 1)
	customer.SelectToBasket(part2Book, 1)

	currentBasket := customer.GetBasket()
	assert.Equal(t, 2, currentBasket.CountBook())
	assert.Equal(t, Euro(16.00), currentBasket.GetTotalPrice())
	assert.Equal(t, Euro(0.8), currentBasket.GetDiscountPrice())
	assert.Equal(t, Euro(15.20), currentBasket.GetNetPrice())

}
