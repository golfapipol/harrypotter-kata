package potter

type Euro float64

func (e Euro) GetPrice() float64 {
	return float64(e)
}
