package task1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicCase(t *testing.T) {
	var input *TaskInput = &TaskInput{
		N: 2,
		M: 3,
		OffersByUnit: []Offers{
			Offers{
				[]Offer{
					Offer{"offer1", 10846332, 1490},
					Offer{"offer2", 682644, 499},
				},
			},
			Offers{
				[]Offer{
					Offer{"offer3", 832784, 14000},
					Offer{"offer4", 3234, 100},
				},
			},
		},
	}

	expected := &Offers{
		[]Offer{
			Offer{"offer3", 832784, 14000},
			Offer{"offer1", 10846332, 1490},
			Offer{"offer2", 682644, 499},
		},
	}

	actual, err := Solution(input)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
	assertEqualOffersArray(t, expected.Offers, actual.Offers)

}

func assertEqualOffers(t *testing.T, expected Offer, actual Offer) {
	assert.Equal(t, expected.OfferId, actual.OfferId)
	assert.Equal(t, expected.MarketSku, actual.MarketSku)
	assert.Equal(t, expected.Price, actual.Price)
}

func assertEqualOffersArray(t *testing.T, expected []Offer, actual []Offer) {
	assert.Equal(t, len(expected), len(actual))

	if len(expected) != len(actual) {
		return
	}

	for i := range expected {
		assertEqualOffers(t, expected[i], actual[i])
	}
}
