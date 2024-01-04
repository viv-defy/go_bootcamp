package basic_number_filtering

import (
	"sort"
	"testing"
)

func areSlicesEqual(got, expected []int) bool {
	if len(got) != len(expected) {
		return false
	}

	sort.Ints(got)
	sort.Ints(expected)
	for i, v := range expected {
		if got[i] != v {
			return false
		}
	}

	return true
}

func TestStory1(t *testing.T) {
	res := EvenNumbers([]int{0, 1, 2, 3, 4, 5, 6, 7, 8})
	expected := []int{0, 2, 4, 6, 8}

	if !areSlicesEqual(res, expected) {
		t.Errorf("expected - %v, got - %v", expected, res)
	}
}

func TestStory2(t *testing.T) {
	res := OddNumbers([]int{0, 1, 2, 3, 4, 5, 6, 7, 8})
	expected := []int{1, 3, 5, 7}

	if !areSlicesEqual(res, expected) {
		t.Errorf("expected - %v, got - %v", expected, res)
	}
}

func TestStory3(t *testing.T) {
	res := PrimeNumbers([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13})
	expected := []int{2, 3, 5, 7, 11, 13}

	if !areSlicesEqual(res, expected) {
		t.Errorf("expected - %v, got - %v", expected, res)
	}
}

func TestStory4(t *testing.T) {
	res := OddAndPrimeNumbers([]int{1, 2, 3, 4, 5, 6, 7, 8})
	expected := []int{3, 5, 7}

	if !areSlicesEqual(res, expected) {
		t.Errorf("expected - %v, got - %v", expected, res)
	}
}

func TestStory5(t *testing.T) {
	res := IsEvenAndMultipleOf5([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20})
	expected := []int{10, 20}

	if !areSlicesEqual(res, expected) {
		t.Errorf("expected - %v, got - %v", expected, res)
	}
}

func TestStory6(t *testing.T) {
	res := IsOddAndMultipleOf3GreaterThan10([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20})
	expected := []int{15}

	if !areSlicesEqual(res, expected) {
		t.Errorf("expected - %v, got - %v", expected, res)
	}
}

func TestStory7(t *testing.T) {
	conditions := []func(int) bool{isOdd, isGreaterThan(5), isMultipleOf(3)}
	res := All([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, conditions)
	expected := []int{9, 15}

	if !areSlicesEqual(res, expected) {
		t.Errorf("expected - %v, got - %v", expected, res)
	}

	conditions = []func(int) bool{isEven, isLessThan(15), isMultipleOf(3)}
	res = All([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, conditions)
	expected = []int{6, 12}

	if !areSlicesEqual(res, expected) {
		t.Errorf("expected - %v, got - %v", expected, res)
	}
}

func TestStory8(t *testing.T) {
	conditions := []func(int) bool{isPrime, isGreaterThan(15), isMultipleOf(5)}
	res := Any([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, conditions)
	expected := []int{2, 3, 5, 7, 10, 11, 13, 15, 16, 17, 18, 19, 20}

	if !areSlicesEqual(res, expected) {
		t.Errorf("expected - %v, got - %v", expected, res)
	}

	conditions = []func(int) bool{isLessThan(6), isMultipleOf(3)}
	res = Any([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, conditions)
	expected = []int{1, 2, 3, 4, 5, 6, 9, 12, 15, 18}

	if !areSlicesEqual(res, expected) {
		t.Errorf("expected - %v, got - %v", expected, res)
	}
}
