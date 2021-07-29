package silvercipher

// Developed by Tharouet Maamouri
// Open source
func FindProgression(d []int, operators string) (r map[int][]int) {
	// default settings
	if operators == "" {
		operators = "asc"
	}
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
	}
	return r
}
