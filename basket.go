package potter

type Customer struct {
	basket Basket
}

func NewCustomer() Customer {
	return Customer{
		basket: NewBasket(),
	}
}

func (c *Customer) SelectToBasket(book Book, quantity int) {
	c.basket.AddBookWithQuantity(book, quantity)
}

func (c Customer) GetBasket() Basket {
	return c.basket
}

type Basket struct {
	books         []SelectedBook
	totalPrice    Euro
	discountPrice Euro
	netPrice      Euro
}

func NewBasket() Basket {
	return Basket{
		books: []SelectedBook{},
	}
}

func (b *Basket) AddBookWithQuantity(selectedBook Book, quantity int) {
	b.books = append(b.books, NewSelectedBook(selectedBook.GetTitle(), selectedBook.GetPrice(), quantity))
}

func (b Basket) CountBook() int {
	return len(b.books)
}
func (b Basket) GetTotalPrice() Euro {
	var sum Euro
	for _, item := range b.books {
		sum += item.GetPrice()
	}
	return sum
}
func (b Basket) GetDiscountPrice() Euro {
	if len(b.books) == 2 {
		return b.GetTotalPrice() * 0.05
	}
	return 0.00
}
func (b Basket) GetNetPrice() Euro {
	return b.GetTotalPrice() - b.GetDiscountPrice()
}

type SelectedBook struct {
	title    string
	isbn     string
	price    Euro
	stock    int
	quantity int
}

func NewSelectedBook(title string, price Euro, quantity int) SelectedBook {
	return SelectedBook{
		title:    title,
		price:    price,
		quantity: quantity,
	}
}

func (b SelectedBook) GetPrice() Euro {
	return b.price
}
