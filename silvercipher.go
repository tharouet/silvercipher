// Developed by Tharouet Maamouri
// Open source
// Package silver cipher provides simple data analytic tools and pattern recognition
package silvercipher

import "fmt"

// Find Progression find int number progression either on an ascending or descending basis
// and return a map of slices with the results
// operator can either be "desc" or "asc"
// minimum is an int and it the minimum number of numbers progressing int he slice
func FindProgression(d []int, operator string, minimum int) (r map[int][]int, err error) {

	// set parameters
	r = make(map[int][]int)
	f := make([]int, 0)
	i := 0

	// minium number of ints in a progression
	if minimum < 1 {
		err = fmt.Errorf("invalid minimum input: %v ! minimum value must but 1 or over ", minimum)
		return nil, err
	}

	minimum = minimum - 1
	f = append(f, d[0])
	if operator == "asc" {
		for k := 1; k < len(d); k++ {
			if d[k-1] < d[k] {
				f = append(f, d[k])
			} else {
				if len(f) > minimum {
					r[i] = f
					i++
				}
				f = []int{}
				f = append(f, d[k])
			}
		}
		if len(f) > minimum {
			r[i] = f
		}
	} else if operator == "desc" {
		for k := 1; k < len(d); k++ {
			if d[k-1] > d[k] {
				f = append(f, d[k])
			} else {
				if len(f) > minimum {
					r[i] = f
					i++
				}
				f = []int{}
				f = append(f, d[k])
			}
		}
		if len(f) > minimum {
			r[i] = f
		}
	} else {
		err = fmt.Errorf("invalid operator input: %v ! operator type can only be desc or asc", operator)
		return r, err
	}
	return r, nil
}
