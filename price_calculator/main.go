package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"de.anikate/price_calculator/tax"
)

func main() {
	taxRates := [3]float64{0, 10, 20}
	fmt.Println("Tax Rates:", taxRates)

	prices, err := readPrices()
	if err != nil {
		return
	}
	fmt.Println("Read prices:", prices)

	taxes := make([]tax.Tax, len(taxRates))
	for index, taxRate := range taxRates {
		taxes[index] = tax.New(taxRate, prices)
		taxes[index].Display()
	}

	err = saveTaxes(taxes)
	if err != nil {
		fmt.Print(err)
		return
	}
}

func readPrices() ([]float64, error) {
	data, err := os.ReadFile("prices.txt")
	if err != nil {
		fmt.Println("Unable to read file.")
		return nil, err
	}

	pricesString := string(data)

	pricesList := strings.Split(pricesString, "\r\n")

	prices := make([]float64, len(pricesList))
	for index, value := range pricesList {
		prices[index], err = strconv.ParseFloat(value, 64)

		if err != nil {
			fmt.Println("Unable to parse prices from file.")
			return nil, err
		}

	}

	return prices, nil
}

func saveTaxes(taxes []tax.Tax) error {

	for _, value := range taxes {
		jsonData, err := json.MarshalIndent(value, "", "\t")
		if err != nil {
			fmt.Print("Unable to convert to JSON data.")
			return err
		}
		os.WriteFile(fmt.Sprintf("tax_%.0f.json", value.TaxRate), jsonData, 0644)

	}
	return nil
}
