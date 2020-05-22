package main

func leetcode1013(A []int) bool {
	sum := make(map[int][]int)
	var tmp int
	for k,v := range A {
		tmp += v
		sum[tmp] = append(sum[tmp], k)
	}
	tmpSum := tmp/3
	if tmpSum*3 != tmp {
		return false
	}
	if tmpSum == 0 {
		if o, ok := sum[0]; ok {
			if len(o) < 3 {
				return false
			}
		}
	}
	//fmt.Println(tmp, tmpSum)
	//fmt.Println(sum)
	indexI,ok := sum[tmpSum]
	if ok {
		indexJ,ok := sum[2*tmpSum]
		if ok {
			for _,v := range indexJ {
				if v > indexI[0] {
					return true
				}
			}
		}
	}
	return false
}

func leetcode1071(str1, str2 string) string {
	if str1 + str2 != str2 + str1 {
		return ""
	}
	return str1[0:GCD(len(str1), len(str2))]
}
