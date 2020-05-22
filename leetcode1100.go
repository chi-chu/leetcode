package main

import(
	"math"
)
func leetcode1103(candies int, num_people int) []int { //10,3 [5,2,3
	d := make([]int, num_people, num_people)
	//how many people received complete gifts
	p := int(math.Sqrt(2 * float64(candies) + 0.25) - 0.5)
	remaining := int(float64(candies) - float64((p + 1)*p) * 0.5)
	rows := p/num_people
	cols := p - rows*num_people
	for k,_ := range d {
		//complete rows
		d[k] = (k+1)*rows + int(float64(rows*(rows-1))*0.5)*num_people
		//cols in the last row
		if k < cols {
			d[k] += k + 1 + rows*num_people
		}
	}
	//remaining candies
	d[cols] += remaining
	return d
}