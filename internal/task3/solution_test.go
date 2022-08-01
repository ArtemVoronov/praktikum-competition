package task3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicCase1(t *testing.T) {
	input := []string{
		"1\tset\t4\ta\t0",
		"2\tset\t3\ta\t1",
		"3\tget\ta",
		"4\tget\ta",
		"-1",
	}

	expected := "false\n" + "false\n" + "true\t1\n" + "true\t0\n"

	actual, err := Solution(input)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestBasicCase2(t *testing.T) {
	input := []string{
		"0\tset\t3\t1\ta",
		"2\tset\t3\t1\tb",
		"3\tget\t1",
		"4\tcancel\t1",
		"-1",
	}

	expected := "false\n" + "false\n" + "true\tb\n" + "false\n"

	actual, err := Solution(input)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestBasicCase3(t *testing.T) {
	input := []string{
		"1\tset\t10\ta\tb",
		"2\tcancel\t8",
		"3\tcancel\t1",
		"4\tcancel\t1",
		"8\tset\t9\tb\ta",
		"9\tget\tb",
		"11\tget\ta",
		"12\tset\t13\ta\tc",
		"13\tcancel\t12",
		"14\tget\ta",
		"-1",
	}

	expected := "false\n" + "false\n" + "true\n" + "false\n" + "false\n" + "true\ta\n" + "false\n" + "false\n" + "false\n" + "true\tc\n"

	actual, err := Solution(input)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}
