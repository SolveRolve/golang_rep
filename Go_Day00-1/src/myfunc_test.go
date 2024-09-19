package main

import (
	"math"
	"testing"
)

const (
	success = "\u2713"
	failed  = "\u2717"
)

func TestData_Mean(t *testing.T) {
	d := NewData()
	t.Log("\tdifferent array for func mean()")
	{
		testID := 0
		t.Logf("\tTest %d:\t", testID)
		{
			d.Numbers = []int{1, 2, 3, 4, 5, 6, 7}
			result := d.Mean()
			exactResult := 4.0
			ok := (result == exactResult)

			if !ok {
				t.Fatalf("\t%s\texact result =%v, function result = %v.", failed, exactResult, result)
			}
			t.Logf("\t%s\t exactResult == result (%v==%v)", success, exactResult, result)
		}

		testID++
		t.Logf("\tTest %d:\t", testID)
		{
			d.Numbers = []int{}
			result := d.Mean()
			exactResult := 0.0
			ok := (result == exactResult)

			if !ok {
				t.Fatalf("\t%s\texact result =%v, function result = %v.", failed, exactResult, result)
			}
			t.Logf("\t%s\t exactResult == result (%v==%v)", success, exactResult, result)
		}
		testID++
		t.Logf("\tTest %d:\t", testID)
		{
			d.Numbers = []int{-1, -2, -3, -5, 100}
			result := d.Mean()
			exactResult := 17.8
			ok := (result == exactResult)

			if !ok {
				t.Fatalf("\t%s\texact result =%v, function result = %v.", failed, exactResult, result)
			}
			t.Logf("\t%s\t exactResult == result (%v==%v)", success, exactResult, result)
		}
	}

}
func TestData_Median(t *testing.T) {
	d := NewData()
	t.Log("\tdifferent array for func median()")
	{
		testID := 0
		t.Logf("\tTest %d:\t", testID)
		{
			d.Numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
			result := d.Median()
			exactResult := 5.0
			ok := (result == exactResult)

			if !ok {
				t.Fatalf("\t%s\texact result =%v, function result = %v.", failed, exactResult, result)
			}
			t.Logf("\t%s\t exactResult == result (%v==%v)", success, exactResult, result)
		}
		testID++
		t.Logf("\tTest %d:\t", testID)
		{
			d.Numbers = []int{}
			result := d.Median()
			exactResult := 0.0
			ok := (result == exactResult)

			if !ok {
				t.Fatalf("\t%s\texact result =%v, function result = %v.", failed, exactResult, result)
			}
			t.Logf("\t%s\t exactResult == result (%v==%v)", success, exactResult, result)
		}
		testID++
		t.Logf("\tTest %d:\t", testID)
		{
			d.Numbers = []int{3, 6, 2, 7, 4}
			result := d.Median()
			exactResult := 4.0
			ok := (result == exactResult)

			if !ok {
				t.Fatalf("\t%s\texact result =%v, function result = %v.", failed, exactResult, result)
			}
			t.Logf("\t%s\t exactResult == result (%v==%v)", success, exactResult, result)
		}
		testID++
		t.Logf("\tTest %d:\t", testID)
		{
			d.Numbers = []int{1, 2}
			result := d.Median()
			exactResult := 1.5
			ok := (result == exactResult)

			if !ok {
				t.Fatalf("\t%s\texact result =%v, function result = %v.", failed, exactResult, result)
			}
			t.Logf("\t%s\t exactResult == result (%v==%v)", success, exactResult, result)
		}
		testID++
		t.Logf("\tTest %d:\t", testID)
		{
			d.Numbers = []int{0, 0}
			result := d.Median()
			exactResult := 0.0
			ok := (result == exactResult)

			if !ok {
				t.Fatalf("\t%s\texact result =%v, function result = %v.", failed, exactResult, result)
			}
			t.Logf("\t%s\t exactResult == result (%v==%v)", success, exactResult, result)
		}
		testID++
		t.Logf("\tTest %d:\t", testID)
		{
			d.Numbers = []int{100}
			result := d.Median()
			exactResult := 100.0
			ok := (result == exactResult)

			if !ok {
				t.Fatalf("\t%s\texact result =%v, function result = %v.", failed, exactResult, result)
			}
			t.Logf("\t%s\t exactResult == result (%v==%v)", success, exactResult, result)
		}
	}
}
func TestData_Mode(t *testing.T) {
	d := NewData()
	t.Log("\tdifferent array for func mode()")
	{
		testID := 0
		t.Logf("\tTest %d:\t", testID)
		{
			d.Numbers = []int{}
			result := d.Mode()
			exactResult := 0
			ok := (result == exactResult)
			if !ok {
				t.Fatalf("\t%s\texact result =%v, function result = %v.", failed, exactResult, result)
			}
			t.Logf("\t%s\t exactResult == result (%v==%v)", success, exactResult, result)
		}
		testID++
		t.Logf("\tTest %d:\t", testID)
		{
			d.Numbers = []int{1}
			d.NumberMap[1] = 1
			result := d.Mode()
			exactResult := 1
			ok := (result == exactResult)

			if !ok {
				t.Fatalf("\t%s\texact result =%v, function result = %v.", failed, exactResult, result)
			}
			t.Logf("\t%s\t exactResult == result (%v==%v)", success, exactResult, result)
		}
		testID++
		t.Logf("\tTest %d:\t", testID)
		{
			d.Numbers = []int{1, 2}
			d.NumberMap[1] = 1
			d.NumberMap[2] = 1
			result := d.Mode()
			exactResult := 1
			ok := (result == exactResult)

			if !ok {
				t.Fatalf("\t%s\texact result =%v, function result = %v.", failed, exactResult, result)
			}
			t.Logf("\t%s\t exactResult == result (%v==%v)", success, exactResult, result)
		}
	}
}
func TestData_StandDispersion(t *testing.T) {
	d := NewData()
	t.Log("\tdifferent array for func mode()")
	{
		testID := 0
		t.Logf("\tTest %d:\t", testID)
		{
			d.Numbers = []int{}
			result := d.StandDispersion()
			exactResult := 0.0
			ok := (result == exactResult)
			if !ok {
				t.Fatalf("\t%s\texact result =%v, function result = %v.", failed, exactResult, result)
			}
			t.Logf("\t%s\t exactResult == result (%v==%v)", success, exactResult, result)
		}
		testID++
		t.Logf("\tTest %d:\t", testID)
		{
			d.Numbers = []int{1}
			result := d.StandDispersion()
			exactResult := 0.0
			ok := (result == exactResult)

			if !ok {
				t.Fatalf("\t%s\texact result =%v, function result = %v.", failed, exactResult, result)
			}
			t.Logf("\t%s\t exactResult == result (%v==%v)", success, exactResult, result)
		}
		testID++
		t.Logf("\tTest %d:\t", testID)
		{
			d.Numbers = []int{23, 34, 6, 234, 4, 6}
			result := d.StandDispersion()
			exactResult := 90.3558
			ok := math.Abs(result-exactResult) < 1e-3

			if !ok {
				t.Fatalf("\t%s\texact result =%v, function result = %v.", failed, exactResult, result)
			}
			t.Logf("\t%s\t exactResult == result (%v==%v)", success, exactResult, result)
		}
	}
}
