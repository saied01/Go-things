package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"net/http"

	"github.com/charmbracelet/huh"
)

var (
	currFrom string
	currTo   string
	amount   string
	result   float64
)

func mainform(curs []huh.Option[string]) *huh.Form {

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Convert from:").
				Options(curs...).
				Value(&currFrom),
		),
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Convert to:").
				Options(curs...).
				Value(&currTo),
		),
		huh.NewGroup(
			huh.NewInput().
				Title("Enter currency amount:").
				Value(&amount).
				Validate(func(val string) error {
					if _, err := strconv.ParseFloat(val, 64); err != nil {
						return errors.New("please enter a valid number")
					}
					return nil
				}),
		),
	)

	return form
}

func getCurrencies(apikey string) (map[string]string, error) {
	url := fmt.Sprintf("https://api.freecurrencyapi.com/v1/currencies?apikey=%s", apikey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var currJson struct {
		Data map[string]struct {
			Name string `json:"name"`
		} `json:"data"`
	}

	err = json.NewDecoder(resp.Body).Decode(&currJson)
	if err != nil {
		return nil, err
	}

	currencyMap := make(map[string]string)
	for code, info := range currJson.Data {
		currencyMap[code] = info.Name
	}

	return currencyMap, nil
}

func main() {

	apikey := os.Getenv("FREECURRENCY_API_KEY")
	currencyMap, err := getCurrencies(apikey)

	var currencies []huh.Option[string]
	for code, name := range currencyMap {
		currencies = append(currencies, huh.NewOption(fmt.Sprintf("%s - %s", code, name), code))
	}

	form := mainform(currencies)
	err = form.Run()
	if err != nil {
		log.Fatal(err)
	}

	url := fmt.Sprintf("https://api.freecurrencyapi.com/v1/latest?apikey=%s&base_currency=%s&currencies=%s", apikey, currFrom, currTo)

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	type Response struct {
		Data map[string]float64 `json:"data"`
	}

	var res Response

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		panic(err)
	}

	rate := res.Data[currTo]

	amountFloat, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		panic(err)
	}

	result = amountFloat * rate

	fmt.Println("conversion: ", result)
}
