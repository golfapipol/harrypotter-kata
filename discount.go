package potter

type Discount struct {
	books      []SelectedBook
	totalPrice Euro
}

func NewDiscount(books []SelectedBook, totalPrice Euro) Discount {
	return Discount{
		books:      books,
		totalPrice: totalPrice,
	}
}

func (d Discount) GetPrice() Euro {
	if len(d.books) == 2 {
		return d.totalPrice * 0.05
	}
	return 0.00
}
