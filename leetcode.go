package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func leetcode6(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	retList := make([]string, numRows, numRows)
	index := 0
	var flag bool
	for _,v := range s {
		if index == 0 {
			flag = true
		}
		if index + 1 == numRows {
			flag = false
		}
		retList[index] += string(v)
		if flag {
			index++
		}else{
			index--
		}
	}
	return strings.Join(retList,"")
}

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

func leetcode22(n int) []string {
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

func leetcode23(lists []*Node) *Node {
	if len(lists) == 0 {
		return nil
	}
	var head, ret *Node
	var tmp int
	for {
		tmp = -1
		for k,v := range lists {
			if v == nil {
				continue
			}
			if tmp == -1 {
				tmp = k
			}else{
				if v.Value < lists[tmp].Value {
					tmp = k
				}
			}
		}
		if tmp == -1 {
			break
		}
		if head == nil{
			head = lists[tmp]
			ret = head
		}else{
			head.Next = lists[tmp]
			head = head.Next
		}
		lists[tmp] = lists[tmp].Next
	}
	return ret
}

func leetcode25(head *Node, k int) *Node{
	var ret, pre, next, current, listC *Node
	var num int
	var flag bool
	pre = nil
	current = head
	for current.Next != nil {
		if num == 0 {
			listC = current
		}
		next = current.Next
		current.Next = pre
		pre = current
		current = next
		num++
		if num == k {
			if !flag {
				ret = pre
				flag = true
			}
			num = 0
			listC.Next = next
		}
	}
	if num != k {
		//反转回去最后不足k个节点
		for listC.Next != nil {

		}
	}
	if ret == nil {
		return head
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

func leetcode39(candidates []int, target int) [][]int {
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
	huisu(nums, res, &ans)
	return ans
}

func huisu(nums []int,res []int, ans *[][]int){
	if len(nums) < 1{
		*ans = append(*ans, res)
		return
	}
	for i := 0; i < len(nums); i++ {
		n := make([]int, len(nums))
		copy(n,nums)//复制一份nums，然后append（n[:i],n[i+1]...），如果用nums执行上述操作，会导致影响nums，不是原有的nums来执行for循环
		huisu(append(n[:i], n[i+1:]...), append(res, nums[i]), ans)
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

//穷举超时。。。。
/*func leetcode62(m, n int) int {
	var ret int
	getAllRoute(m, n, 1, 1, &ret)
	return ret
}
func getAllRoute(m, n, right, down int, ret *int) {
	if right == m && down == n {
		*ret += 1
		return
	}
	if right < m {
		getAllRoute(m, n, right+1, down, ret)
	}
	if down < n {
		getAllRoute(m, n, right, down+1, ret)
	}
}*/

func leetcode62(m, n int) int {
	//C(m+n-2, m-1)
	return 0
}

//又超时 我擦。。。
/*func leetcode64(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	var ret int
	ret = math.MaxInt64
	getShortRoute(grid, 0, 0, 0, &ret)
	return ret
}
func getShortRoute(grid [][]int, right, down, tmpVal int, minRounte *int) {
	tmpVal += grid[down][right]
	if right + 1 == len(grid[0]) && down + 1 == len(grid) {
		if tmpVal < *minRounte {
			*minRounte = tmpVal
		}
		return
	}
	if right + 1 < len(grid[0]) {
		getShortRoute(grid, right + 1, down, tmpVal, minRounte)
	}
	if down + 1 < len(grid) {
		getShortRoute(grid, right, down + 1, tmpVal, minRounte)
	}
}*/
func leetcode64(grid [][]int) int {
	//动态规划
	if len(grid) == 0 {
		return 0
	}
	var mirrorGrid [][]int
	//工程级代码 一般不允许修改原变量
	lenx := len(grid[0])
	leny := len(grid)
	mirrorGrid = make([][]int, leny)
	copy(mirrorGrid, grid)
	//fmt.Println(mirrorGrid)
	for i := lenx - 1; i >=0; i-- {
		for j:= leny - 1; j >= 0; j-- {
			if i + 1 <= lenx - 1 && j + 1 <= leny - 1 {
				if mirrorGrid[j+1][i] > mirrorGrid[j][i+1] {
					mirrorGrid[j][i] = mirrorGrid[j][i+1] + grid[j][i]
				}else{
					mirrorGrid[j][i] = mirrorGrid[j+1][i] + grid[j][i]
				}
			}else if i + 1 > lenx - 1 && j + 1 <= leny - 1 {
				mirrorGrid[j][i] = mirrorGrid[j+1][i] + grid[j][i]
			}else if i + 1 <= lenx - 1 && j + 1 > leny - 1 {
				mirrorGrid[j][i] = mirrorGrid[j][i+1] + grid[j][i]
			}else {
				mirrorGrid[j][i] = grid[j][i]
			}
		}
	}
	return mirrorGrid[0][0]
}

func leetcode69(x int) int {
	//最快应该用二分。
	if x == 0 {
		return 0
	}
	if x < 4 {
		return 1
	}
	var ret int
	main := 2
	for main*main <= x {
		ret = main
		main *= main
	}
	for ret*ret <= x {
		ret++
	}
	return ret - 1
}

//上一波上下边界要在 搜索二叉树范围内
func leetcode98(root *TreeNode) bool {
	return checkBST(root, math.MinInt64, math.MaxInt64)
}

func checkBST(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return checkBST(root.Left, lower, root.Val) && checkBST(root.Right, root.Val, upper)
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