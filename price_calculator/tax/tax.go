package tax

import "fmt"

type Tax struct {
	TaxRate     float64
	Prices      []float64
	TaxedPrices []float64
}

func New(taxRate float64, prices []float64) Tax {

	taxedPrices := make([]float64, len(prices))
	for index, price := range prices {
		taxedPrices[index] = price + (price * taxRate / 100)
	}

	return Tax{
		taxRate,
		prices,
		taxedPrices,
	}
}

func (t Tax) Display() {
	fmt.Printf("Tax Rate: %.2f%%\n", t.TaxRate)
	for index, tPrice := range t.TaxedPrices {
		fmt.Printf("\t%.2f -> %.2f\n", t.Prices[index], tPrice)
	}
	fmt.Println()
}
