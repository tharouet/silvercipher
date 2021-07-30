# silvercipher
Free Data Analytics Tools Package


<h1>Installation:</h1>
	
go get "github.com/tharouet/silvercipher"

go install "github.com/tharouet/silvercipher"


<h1>Example:</h1>
<h2>Code: main.go </h2>
<div>


	s := []int{23, 4, 5, 6, 87, 9, 5, 23, 3, 4, 5, 6, 12, 34, 76, 54, 343, 2, 4, 33, 37, 43}
	results, err := silvercipher.FindProgression(s, "asc")
	if err != nil {
		fmt.Println(err)
	} else {
		for _, v := range results {
			fmt.Println(v)
		}
	}


<h2>Result:</h2>

go run main.go 

[2 4 33 37 43]
[23]
[4 5 6 87]
[9]
[5 23]
[3 4 5 6 12 34 76]
[54 343]
