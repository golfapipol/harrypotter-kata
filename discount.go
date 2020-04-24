package potter

type Discount struct {
	books      []SelectedBook
	totalPrice Euro
}

type SetOfBook struct {
	books map[string]SelectedBook
}

func NewSetOfBook() SetOfBook {
	return SetOfBook{
		books: map[string]SelectedBook{},
	}
}

func (s *SetOfBook) IsBookExist(book SelectedBook) bool {
	_, existed := s.books[book.GetTitle()]
	return existed
}

func (s *SetOfBook) AddBook(book SelectedBook) {
	s.books[book.GetTitle()] = book
}

func (s SetOfBook) GetPrice() Euro {
	var sum Euro
	for _, book := range s.books {
		sum += book.GetPrice()
	}
	if len(s.books) == 2 {
		return sum * 0.05
	}
	return 0.00
}

func NewDiscount(books []SelectedBook, totalPrice Euro) Discount {
	return Discount{
		books:      books,
		totalPrice: totalPrice,
	}
}

func (d Discount) GetPrice() Euro {
	sets := []SetOfBook{}
	for _, book := range d.books {
		if len(sets) < book.GetQuantity() {
			setsLength := len(sets)
			for i := 0; i <= book.GetQuantity()-setsLength; i++ {
				sets = append(sets, NewSetOfBook())
			}
		}
		for i, currentSet := 0, 0; i < book.GetQuantity(); {
			if sets[currentSet].IsBookExist(book) {
				currentSet++
				continue
			}
			sets[currentSet].AddBook(book)
			currentSet++
			i++
		}
	}
	var sum Euro
	for _, setBook := range sets {
		sum += setBook.GetPrice()
	}
	return sum
}
