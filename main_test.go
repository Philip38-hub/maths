package mathskill

import (
	"testing"
	"reflect"
 )

func TestSort(t *testing.T) {
	want := []float64{1, 2, 3, 4}
	got := Sort([]float64{3, 4, 1, 2})

	if !reflect.DeepEqual(want, got){
		t.Errorf("got %v want %v", got, want)
	}
}

func TestAverage(t *testing.T) {
	want := 6.333333333333333
	got := Average([]float64{6, 7, 9, 8, 3, 5})

	if want != got{
		t.Errorf("got %v want %v", got, want)
	}
}

func TestStandardDev(t *testing.T) {
	want := 1.9720265943665387
	got := StandardDev([]float64{6, 7, 9, 8, 3, 5})

	if want != got{
		t.Errorf("got %v want %v", got, want)
	}
}

func TestVariance(t *testing.T) {
	want := 3.8888888888888893 
	got := Variance([]float64{6, 7, 9, 8, 3, 5})

	if want != got{
		t.Errorf("got %v want %v", got, want)
	}
}

func TestMedian(t *testing.T) {
	want := 6.5
	got := Median([]float64{6, 7, 9, 8, 3, 5})

	if want != got{
		t.Errorf("got %v want %v", got, want)
	}
}



