package potter

type Book struct {
	title string
	isbn  string
	price Euro
	stock int
}

func NewBook(title string, price Euro) Book {
	return Book{
		title: title,
		price: price,
	}
}

func (b Book) GetPrice() Euro {
	return b.price
}

func (b Book) GetTitle() string {
	return b.title
}
