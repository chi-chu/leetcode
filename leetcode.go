package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getRomaNum(n int) string {
	ret := ""
	for n > 0 {
		if n >= 1000 {
			ret += "M"
			n -= 1000
		}else if n >= 500 {
			ret += "D"
			n -= 500
		}else if n >= 400 {
			ret += "CD"
			n -= 400
		}else if n >=300 {
			ret += "CCC"
			n -= 300
		}else if n >= 200 {
			ret += "CC"
			n -= 200
		}else if n >= 100 {
			ret += "C"
			n -= 100
		}else if n >= 50 {
			ret += "L"
			n -= 50
		}
	}
	return ret
}

func checkstring(s string, p string) bool {
	//* 42  .46
	sl := len(s)
	pl := len(p)
	i,j := 0,0
	var pre byte
	for i<sl {
		if s[i] == p[j] {
			pre = s[i]
			i += 1
			j += 1
		}else if p[j] == '.' {
			pre = '.'
			i += 1
			j += 1
		}else if p[j] == '*' {
			if pre == s[i] {
				i += 1
			}else if pre == '.' {
				i += 1
			}else{
				return false
			}
		}else{
			return false
		}
		if j >= pl {
			return false
		}
	}
	if j != 0 && j != pl-1 {
		return false
	}
	return true
}

func threeSum(nums []int) [][]int {
	var ret [][]int
	quickSort(nums, 0, len(nums) -1)
	fmt.Println(nums)
	return ret
}

//leetcode 22
func generateParenthesis(n int) []string {
	ret := make([]string,0)
	var leftkh,rightkh int
	var tmp string
	generate(2*n, n, leftkh, rightkh, tmp, &ret)
	return ret
}

func generate(n, total, leftkh, rightkh int, tmp string, ret *[]string) {
	if leftkh+rightkh == 2*total {
		*ret = append(*ret, tmp)
		return
	}
	for n > 0 {
		if rightkh == leftkh {
			generate(n-1, total, leftkh+1, rightkh, tmp+"(", ret)
		}else if leftkh == total {
			generate(n-1, total, leftkh, rightkh+1, tmp+")", ret)
		}else{
			generate(n-1, total, leftkh+1, rightkh, tmp+"(", ret)
			generate(n-1, total, leftkh, rightkh+1, tmp+")", ret)
		}
		n--
	}
}

//leetcode 6
func zstring(s string, numRows int) string {
	ret := ""
	index := 0
	lens := len(s)
	for i:=0; i<numRows; i++ {
		for index < lens {
			
		}
	}
	return ret
}

//leetcode32
func longestValidParentheses(s string) int {
	var tmplen, maxlen, k int
	var simg string
	for _,v := range s {
		if v == '(' {
			simg += "("
			k++
		}else{
			if k < 1 {
				continue
			}
			if simg[k-1] == '(' {
				tmplen += 2
				simg = simg[:k-1]
				k--
				if tmplen > maxlen {
					maxlen = tmplen
				}
				
			}else{
				if tmplen > maxlen {
					maxlen = tmplen
				}
				tmplen = 0
			}
		}
		fmt.Println(tmplen, "---", simg, "---", k)
	}
	if tmplen > maxlen {
		maxlen = tmplen
	}
	return maxlen
}

//leetcode 39
func combinationSum(candidates []int, target int) [][]int {
	var ret [][]int
	for i,_ := range candidates{
		getTarget(candidates, i, target, 0, []int{}, &ret)
	}
	return ret
}

func getTarget(k []int, index int, target int, tmpsum int, tmpList []int, ret *[][]int) {
	tmpsum += k[index]
	tmpList = append(tmpList, k[index])
	if tmpsum == target {
		*ret = append(*ret, tmpList)
		return
	}
	if tmpsum < target {
		for i:=index;i<len(k);i++ {
			tmp := make([]int,len(tmpList))
			copy(tmp, tmpList)
			getTarget(k, i, target, tmpsum, tmp, ret)
		}
	}
}

func leetcode41(nums []int) int {
	var ret int

	return ret
}

func leetcode42(nums []int) int {
	var ret, maxL, lens, index int
	lens = len(nums)
	if lens== 0 {
		return 0
	}
	for k, v := range nums {
		if v > maxL {
			maxL = v
			index = k
		}
	}
	var i, j, tmpheight int
	j = lens-1
	for i<=index {
		i++
		if nums[i] > tmpheight {
			tmpheight = nums[i]
		}else{
			ret += tmpheight - nums[i]
		}
	}
	tmpheight = 0
	for j>=index {
		j--
		if nums[j] > tmpheight {
			tmpheight = nums[j]
		}else{
			ret += tmpheight - nums[j]
		}
	}
	return ret
}

func leetcode43(num1, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	var ret string
	var tmpret []string
	var tmpretlen []int
	len1 := len(num1)
	len2 := len(num2)
	tmpmul := 0
	for i:=len2-1;i>=0;i-- {
		tmpmulstr := ""
		tmpmul = 0
		int2, _ := strconv.Atoi(string(num2[i]))
		for j:=len1-1;j>=0;j-- {
			int1, _ := strconv.Atoi(string(num1[j]))
			tmpmul = tmpmul + int1 * int2
			tmpmulstr = strconv.Itoa(tmpmul%10) + tmpmulstr
			tmpmul /= 10
		}
		if tmpmul != 0 {
			tmpmulstr = strconv.Itoa(tmpmul%10) + tmpmulstr
		}
		tmpretlen = append(tmpretlen, len(tmpmulstr)+len2-i-1)
		tmpret = append(tmpret, tmpmulstr + strings.Repeat("0", len2-i-1))
	}
	fmt.Println(tmpret)
	var flag bool
	var remain int
	var tmpsum int
	k := 0
	for {
		flag = true
		tmpsum = 0
		for index,v := range tmpret {
			if len(v) > k {
				flag = false
				tmp, _ := strconv.Atoi(string(v[tmpretlen[index]-1-k]))
				tmpsum += tmp
			}
		}
		fmt.Println(tmpsum)
		if flag {
			if remain != 0 {
				ret = strconv.Itoa(remain) + ret
			}
			break
		}
		tmpsum += remain
		ret = strconv.Itoa(tmpsum%10) + ret
		remain = tmpsum/10
		k++
	}
	return ret
}

func leetcode45(nums []int)int {
	var ret int
	// 贪心算法
	return ret
}

func leetcode46(nums []int) [][]int {
	var ans [][]int
	var res []int
	huisu(nums,res,&ans)
	return ans
}

func huisu(nums []int,res []int,ans *[][]int){
	if len(nums)<1{
		*ans=append(*ans,res)
		return
	}
	for i:=0;i<len(nums);i++{
		n:=make([]int,len(nums))
		copy(n,nums)//复制一份nums，然后append（n[:i],n[i+1]...），如果用nums执行上述操作，会导致影响nums，不是原有的nums来执行for循环
		huisu(append(n[:i],n[i+1:]...),append(res,nums[i]),ans)
	}
}

func leetcode48(matrix [][]int){
	var lens, tmpVal, changVal int
	lens = len(matrix)
	for j := 0; j < lens/2; j++ {
		for i := j; i < lens-1-j; i++ {
			tmpVal = matrix[i][lens-j-1]
			matrix[i][lens-j-1] = matrix[j][i]
			changVal = matrix[lens-j-1][lens-i-1]
			matrix[lens-j-1][lens-i-1] = tmpVal
			tmpVal = matrix[lens-i-1][j]
			matrix[lens-i-1][j] = changVal
			matrix[j][i] = tmpVal
		}
	}
}

func leetcode49(strs []string) [][]string {
	ret := make([][]string,0,4)
	constMap := map[byte]int{'a':2,'b':3,'c':5,'d':7,'e':11,'f':13,'g':17,'h':19,'i':23,'j':29,'k':31,'l':37,'m':41,
		'n':43,'o':47,'p':53,'q':59,'r':61,'s':67,'t':71,'u':73,'v':79,'w':83,'x':89,'y':97,'z':101}
	sumList := make([]int,0,8)
	for _, v := range strs {
		sumTmp := 1
		for _,word := range v {
			if num, ok := constMap[byte(word)]; ok{
				sumTmp *= num
			}
		}
		var flag bool
		for k, sumExp := range sumList {
			if sumTmp == sumExp {
				flag = true
				ret[k] = append(ret[k], v)
				break
			}
		}
		if !flag {
			sumList = append(sumList, sumTmp)
			ret = append(ret, []string{v})
		}
	}
	return ret
}


func leetcode50(x float64, n int) float64 {
	if n < 1 {
		x = 1/x
		n = -n
	}
	var ret float64 = 1
	exp := n
	for exp != 0{
		tmp, remain := getLimitPow(x, exp)
		ret *= tmp
		exp -= remain
	}
	return ret
}

func getLimitPow(x float64, n int) (float64, int) {
	exp := 1
	for !(exp <= n && 2*exp > n) {
		x *= x
		exp *= 2
	}
	return x, exp
}

func leetcode300(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var ret int
	var tmp []int
	tmp = append(tmp, nums[0])
	for _,v := range nums {
		//fmt.Println(tmp,"---", v)
		if v > tmp[ret] {
			tmp = append(tmp, v)
			ret++
		}else{
			var left, right int
			right = ret
			for left < right {
				mid := (right - left) >>1 + left
				if tmp[mid] < v {
					//只有小于 才 +1
					left = mid + 1
				} else {
					right = mid
				}
			}
			tmp[left] = v
		}
	}
	return ret + 1
}

func leetcode515(root *TreeNode) []int {
	var ret []int
	getLevelMax(&ret, root, 0)
	return ret
}

func getLevelMax(ret *[]int, root *TreeNode, level int) {
	if len(*ret) - 1 < level {
		*ret = append(*ret, root.Val)
	}
	if root.Val > (*ret)[level] {
		(*ret)[level] = root.Val
	}
	if root.Left != nil {
		getLevelMax(ret, root.Left, level+1)
	}
	if root.Right != nil {
		getLevelMax(ret, root.Right, level+1)
	}
}