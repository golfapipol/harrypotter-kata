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
	books []SelectedBook
}

func NewBasket() Basket {
	return Basket{
		books: []SelectedBook{},
	}
}

func (b *Basket) AddBookWithQuantity(selectedBook Book, quantity int) {
	b.books = append(b.books, NewSelectedBook(selectedBook.GetTitle(), selectedBook.GetPrice(), quantity))
}

func (b Basket) GetBooks() []SelectedBook {
	return b.books
}
func (b Basket) CountBook() int {
	return len(b.books)
}
func (b Basket) GetTotalPrice() Euro {
	var sum Euro
	for _, item := range b.books {
		sum += item.GetPrice() * Euro(item.GetQuantity())
	}
	return sum
}
func (b Basket) GetDiscountPrice() Euro {
	discount := NewDiscount(b.books, b.GetTotalPrice())
	return discount.GetPrice()
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

func (b SelectedBook) GetTitle() string {
	return b.title
}
func (b SelectedBook) GetPrice() Euro {
	return b.price
}

func (b SelectedBook) GetQuantity() int {
	return b.quantity
}
