// Developed by Tharouet Maamouri
// Open source
// Package silver cipher provides simple data analytic tools and pattern recognition
package silvercipher

import "fmt"

// Find Progression find int number progression either on an ascending or descending basis
// and return a map of slices with the results
// operators can either be "desc" or "asc"
func FindProgression(d []int, operators string) (r map[int][]int, err error) {
	r = make(map[int][]int)
	f := make([]int, 0)
	i := 0
	f = append(f, d[0])
	if operators == "asc" {
		for k := 1; k < len(d); k++ {
			if d[k-1] < d[k] {
				f = append(f, d[k])
			} else {
				r[i] = f
				i++
				f = []int{}
				f = append(f, d[k])
			}
		}
		r[i] = f
	} else if operators == "desc" {
		for k := 1; k < len(d); k++ {
			if d[k-1] > d[k] {
				f = append(f, d[k])
			} else {
				r[i] = f
				i++
				f = []int{}
				f = append(f, d[k])
			}
		}
		r[i] = f
	} else {
		err = fmt.Errorf("invalid input: <%v> ! operator type can only be desc or asc", operators)
		return r, err
	}
	return r, nil
}
