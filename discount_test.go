package potter_test

import (
	. "potter"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiscount2Part1Book1Part2Book(t *testing.T) {
	part1Book := NewBook("Harry Potter Part 1", Euro(8.00))
	part2Book := NewBook("Harry Potter Part 2", Euro(8.00))
	customer := NewCustomer()

	customer.SelectToBasket(part1Book, 2)
	customer.SelectToBasket(part2Book, 1)

	currentBasket := customer.GetBasket()

	discount := NewDiscount(currentBasket.GetBooks(), currentBasket.GetTotalPrice())

	assert.Equal(t, Euro(0.8), discount.GetPrice())
}
