package task1

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

const MAX_OFFERS_IN_UNIT = 200

type TaskInput struct {
	N            int
	M            int
	OffersByUnit []Offers
}

type Offer struct {
	OfferId   string `json:"offer_id"`
	MarketSku int    `json:"market_sku"`
	Price     int    `json:"price"`
}

type Offers struct {
	Offers []Offer `json:"offers"`
}

type ByPrice []Offer

func (a ByPrice) Len() int           { return len(a) }
func (a ByPrice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPrice) Less(i, j int) bool { return a[i].Price > a[j].Price }

func Solution(input *TaskInput) (*Offers, error) {
	var result *Offers

	var allOffers []Offer = make([]Offer, 0, input.N*MAX_OFFERS_IN_UNIT)
	for _, unitOffers := range input.OffersByUnit {
		allOffers = append(allOffers, unitOffers.Offers...)
	}

	sort.Sort(ByPrice(allOffers))

	result = &Offers{allOffers[:input.M]}

	fmt.Printf("%v\n", result)

	return result, nil
}

func ReadInput(path string) (*TaskInput, error) {
	var result *TaskInput
	data, err := os.ReadFile(path)
	if err != nil {
		return result, err
	}

	input := string(data)

	re := regexp.MustCompile("\n")
	split := re.Split(input, -1)

	var n, m int
	var offersByUnit []Offers
	for i, line := range split {
		if i != 0 {
			var offers Offers
			err = json.Unmarshal([]byte(line), &offers)
			if err != nil {
				return result, err
			}
			offersByUnit = append(offersByUnit, offers)
		} else {
			re2 := regexp.MustCompile(" ")
			digits := re2.Split(line, -1)
			n, err = strconv.Atoi(digits[0])
			if err != nil {
				return result, err
			}
			m, err = strconv.Atoi(digits[1])
			if err != nil {
				return result, err
			}
		}
	}

	return &TaskInput{n, m, offersByUnit}, nil
}

func WriteOutput(path string, offers *Offers) error {
	data, err := json.Marshal((*offers))
	if err != nil {
		return err
	}
	err = os.WriteFile(path, data, 0666)
	return err
}
