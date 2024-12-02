package main

import "testing"

func compare(t *testing.T, want int, actual int) {
	if want != actual {
		t.Fatalf(`Expected %d, got %d`, want, actual)
	}
}

func TestStar1Sample(t *testing.T) {
	want := 2
	actual := Star1(1)
	compare(t, want, actual)
}

func TestStar1(t *testing.T) {
	want := 411
	actual := Star1(0)
	compare(t, want, actual)
}

func TestStar2Sample(t *testing.T) {
	want := 4
	actual := Star2(1)
	compare(t, want, actual)
}

func TestStar2(t *testing.T) {
	want := 465
	actual := Star2(0)
	compare(t, want, actual)
}
