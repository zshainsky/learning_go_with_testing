package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		got := Sum(numbers)
		want := 45

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}

	})
}

func TestSumAll(t *testing.T) {

	t.Run("testing sum of two inputs", func(t *testing.T) {
		num1 := []int{1, 2}
		num2 := []int{0, 9}

		got := SumAll(num1, num2)
		want := []int{3, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v given, SumAll(%v, %v)", got, want, num1, num2)
		}
	})

	t.Run("testing sum of one input", func(t *testing.T) {
		num1 := []int{1, 1, 1}

		got := SumAll(num1)
		want := []int{3}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v given, SumAll(%v)", got, want, num1)
		}
	})

	t.Run("testing sum of zero input", func(t *testing.T) {
		num1 := []int{}

		got := SumAll(num1)
		want := []int{0}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v given, SumAll(%v)", got, want, num1)
		}
	})

}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v given", got, want)
		}
	}

	t.Run("testing sum 2 inputs ", func(t *testing.T) {
		num1 := []int{0, 2}
		num2 := []int{3, 9}

		got := SumAllTails(num1, num2)
		want := []int{2, 9}
		checkSums(t, got, want)

	})
	t.Run("testing sum of empty slice", func(t *testing.T) {
		num1 := []int{}
		num2 := []int{3, 9}

		got := SumAllTails(num1, num2)
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}
