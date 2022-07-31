package task2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicCaseMonth(t *testing.T) {
	inputPeriod := Period{"2022-01-10", "2022-03-25"}
	inputPeriodType := MONTH

	expected := []Period{
		Period{"2022-01-10", "2022-01-31"},
		Period{"2022-02-01", "2022-02-28"},
		Period{"2022-03-01", "2022-03-25"},
	}

	acutal, err := Solution(inputPeriodType, inputPeriod)

	assert.Nil(t, err)
	assertEqualPeriodSlices(t, expected, acutal)
}
func TestBasicCaseWeek(t *testing.T) {
	inputPeriod := Period{"2022-01-30", "2022-03-23"}
	inputPeriodType := WEEK

	expected := []Period{
		Period{"2022-01-30", "2022-01-30"},
		Period{"2022-01-31", "2022-02-06"},
		Period{"2022-02-07", "2022-02-13"},
		Period{"2022-02-14", "2022-02-20"},
		Period{"2022-02-21", "2022-02-27"},
		Period{"2022-02-28", "2022-03-06"},
		Period{"2022-03-07", "2022-03-13"},
		Period{"2022-03-14", "2022-03-20"},
		Period{"2022-03-21", "2022-03-23"},
	}

	acutal, err := Solution(inputPeriodType, inputPeriod)

	assert.Nil(t, err)
	assertEqualPeriodSlices(t, expected, acutal)
}

func TestBasicCaseFriday13(t *testing.T) {
	inputPeriod := Period{"2022-01-12", "2022-05-15"}
	inputPeriodType := FRIDAY_THE_13TH

	expected := []Period{
		Period{"2022-01-12", "2022-05-12"},
		Period{"2022-05-13", "2022-05-15"},
	}

	acutal, err := Solution(inputPeriodType, inputPeriod)

	assert.Nil(t, err)
	assertEqualPeriodSlices(t, expected, acutal)
}

func TestBasicCaseQuarter(t *testing.T) {
	inputPeriod := Period{"2022-01-12", "2022-08-15"}
	inputPeriodType := QUARTER

	expected := []Period{
		Period{"2022-01-12", "2022-03-31"},
		Period{"2022-04-01", "2022-06-30"},
		Period{"2022-07-01", "2022-08-15"},
	}

	actual, err := Solution(inputPeriodType, inputPeriod)

	assert.Nil(t, err)
	assertEqualPeriodSlices(t, expected, actual)
}

func TestBasicCaseYear(t *testing.T) {
	inputPeriod := Period{"2022-01-11", "2025-08-15"}
	inputPeriodType := YEAR

	expected := []Period{
		Period{"2022-01-11", "2022-12-31"},
		Period{"2023-01-01", "2023-12-31"},
		Period{"2024-01-01", "2024-12-31"},
		Period{"2025-01-01", "2025-08-15"},
	}

	acutal, err := Solution(inputPeriodType, inputPeriod)

	assert.Nil(t, err)
	assertEqualPeriodSlices(t, expected, acutal)
}

func assertEqualPeriods(t *testing.T, expected *Period, actual *Period) {
	assert.Equal(t, (*expected).Start, (*actual).Start)
	assert.Equal(t, (*expected).End, (*actual).End)
}

func assertEqualPeriodSlices(t *testing.T, expected []Period, actual []Period) {
	assert.Equal(t, len(expected), len(actual))

	if len(expected) != len(actual) {
		return
	}

	for i := range expected {
		assertEqualPeriods(t, &expected[i], &actual[i])
	}
}
