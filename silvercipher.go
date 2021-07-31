// Developed by Tharouet Maamouri
// Package silvercipher provides simple data analytic tools and pattern recognition
package silvercipher

import "fmt"

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

func GetProgressingNumbersOnly(d []float64, operator string) (f []float64, err error) {

	if operator == "asc" {
		for k := 1; k < len(d); k++ {
			if d[k-1] < d[k] {
				f = append(f, d[k])
			}
		}
	} else if operator == "desc" {
		for k := 1; k < len(d); k++ {
			if d[k-1] > d[k] {
				f = append(f, d[k])
			}
		}
	} else {
		err = fmt.Errorf("invalid operator input: %v - operator type can only be desc or asc", operator)
		return nil, err
	}
	return f, nil
}
