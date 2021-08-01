// Developed by Tharouet Maamouri
// Package silvercipher provides simple data analytic tools and pattern recognition
package silvercipher

import (
	"fmt"
)

// Function to find ascending or descending succession of floats in a slice of integers
// and return a map of slices with the results
// operator: can either be "desc" or "asc"
// minimum: is an int and is the minimum number of floats in a progression
func FindProgressions(d []float64, operator string, minimum int) (r map[int][]float64, err error) {

	r = make(map[int][]float64)
	f := make([]float64, 0)
	i := 0

	if minimum < 1 {
		err = fmt.Errorf("invalid minimum input: %v - minimum value must but 1 or over ", minimum)
		return nil, err
	}

	minimum = minimum - 1
	f = append(f, d[0])
	if operator == "asc" {
		for k := 1; k < len(d); k++ {
			if d[k-1] > d[k] {
				if len(f) > minimum {
					r[i] = f
					i++
				}
				f = []float64{}
			}
			f = append(f, d[k])
		}
		if len(f) > minimum {
			r[i] = f
		}
	} else if operator == "desc" {
		for k := 1; k < len(d); k++ {
			if d[k-1] < d[k] {
				if len(f) > minimum {
					r[i] = f
					i++
				}
				f = []float64{}
			}
			f = append(f, d[k])
		}
		if len(f) > minimum {
			r[i] = f
		}
	} else {
		err = fmt.Errorf("invalid operator input: %v - operator type can only be desc or asc", operator)
		return r, err
	}
	return r, nil
}

// Function to find the number of desc and asc progressions
// and return an int
// operator: can either be "desc" or "asc"
// minimum: is an int and is the minimum number of floats in a progression
func FindNumberOfProgressions(d []float64, operator string, minimum int) (l int, err error) {
	result, err := FindProgressions(d, operator, minimum)
	l = len(result)
	return l, err
}

// Function to get the ascending or descending numbers only in a slice of floats
// and return an int
// operator: can either be "desc" or "asc"
// minimum: is an int and is the minimum number of floats in a progression
func GetProgressingNumbersOnly(d []float64, operator string) (f []float64, err error) {
	z := 0
	f = append(f, d[0])
	if operator == "asc" {
		for k := 1; k < len(d); k++ {
			if d[z] < d[k] {
				f = append(f, d[k])
				z = k
			}
		}
	} else if operator == "desc" {
		for k := 1; k < len(d); k++ {
			if d[z] > d[k] {
				f = append(f, d[k])
				z = k
			}
		}
	} else {
		err = fmt.Errorf("invalid operator input: %v - operator type can only be desc or asc", operator)
		return nil, err
	}
	return f, nil
}

// Take a slice of floats and return the average, highest and lowest values and error if
// slice is empty
func Average(d []float64) (average, highest, lowest float64, err error) {
	l := len(d)
	if l == 0 {
		err = fmt.Errorf("invalid slice input: %v - slice must not be empty", d)
		return 0, 0, 0, err
	}
	var t float64
	highest = d[0]
	lowest = d[0]
	for _, v := range d {
		t = t + v
		if highest > v {
			highest = v
		}
		if lowest < v {
			lowest = v
		}
	}
	average = t / float64(len(d))
	return average, highest, lowest, nil
}

//Insert slice <new> in the n position of slice <original>
func AddToSlice(original, new []float64, n int) (r []float64) {
	a := original[:n]
	b := original[n:]
	r = append(r, a...)
	r = append(r, new...)
	r = append(r, b...)
	return
}

//Remove section from position start to position end from <original> slice
func RemoveFromSlice(original []float64, start, end int) (r []float64) {
	a := original[:start]
	b := original[end:]
	r = append(r, a...)
	r = append(r, b...)
	return r
}

//Invert slice content
func InvertSlice(original []float64) (r []float64) {
	for i := (len(original) - 1); i >= 0; i-- {
		r = append(r, original[i])
	}
	return
}
