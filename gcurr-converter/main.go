package main

import (
	// "encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/charmbracelet/huh"
	// "net/http"
)

var (
	currFrom string
	currTo   string
	amount   string
	result   int
)

func mainform(curs []string) *huh.Form {

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Convert from:").
				Options(
					huh.NewOption(curs[0], curs[0]),
					huh.NewOption(curs[1], curs[1]),
					huh.NewOption(curs[2], curs[2]),
				).
				Value(&currFrom),
		),
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Convert to:").
				Options(
					huh.NewOption(curs[0], curs[0]),
					huh.NewOption(curs[1], curs[1]),
					huh.NewOption(curs[2], curs[2]),
				).
				Value(&currTo),
		),
		huh.NewGroup(
			huh.NewInput().
				Title("Enter currency amount:").
				Value(&amount).
				Validate(func(val string) error {
					if _, err := strconv.Atoi(val); err != nil {
						return errors.New("please enter a valid integer")
					}
					return nil
				}),
		),
	)

	return form
}

func main() {

	currencies := []string{"USD", "ARS", "EUR"}
	form := mainform(currencies)
	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("conversion: %d", result)
	// rateFrom := data.Rates[currencyFrom]
	// rateTo := data.Rates[currencyTo]
	//
	// converted := amount * (rateTo / rateFrom)
}
