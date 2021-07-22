package assignment

import "strconv"

//0. createa a numtype []string
type zeroTen []string

//1. Create Slice of Int
func NewNum() zeroTen {
	oddOrEven := zeroTen{}
	var numStr string

	numLoops := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, nloop := range numLoops {
		if nloop%2 == 0 {
			even := "even"
			numStr = strconv.Itoa(i)
			oddOrEven = append(oddOrEven, numStr+" is "+even)
		} else {
			odd := "odd"
			numStr = strconv.Itoa(i)
			oddOrEven = append(oddOrEven, numStr+" is "+odd)
		}
	}
	return oddOrEven

}
