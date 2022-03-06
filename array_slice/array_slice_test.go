package arrayslice

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got int, want int, numbers []int) {
		t.Helper()
		if got != want {
			t.Errorf("expected %d but got %d. Given %v", want, got, numbers)
		}
	}

	t.Run("collection of any", func(t *testing.T) {
		number := []int{1, 2, 3}

		got := Sum(number)
		want := 6

		assertCorrectMessage(t, got, want, number)
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got []int, want []int) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("Make the sums of slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 4, 1}, []int{0, 9})
		want := []int{7, 9}
		assertCorrectMessage(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		assertCorrectMessage(t, got, want)
	})
}
