package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func leetcode5(s string) string {
	if s == "" {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := expandAroundCenter(s, i, i)
		left2, right2 := expandAroundCenter(s, i, i + 1)
		if right1 - left1 > end - start {
			start, end = left1, right1
		}
		if right2 - left2 > end - start {
			start, end = left2, right2
		}
	}
	return s[start:end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1 , right+1 { }
	return left + 1, right - 1
}

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
	QuickSort(nums, 0, len(nums) -1)
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

//卧槽 巨tm难。。。 时间空间复杂度最低。。
func leetcode25(head *Node, k int) *Node{
	var ret, pre, next, current, listHeadone, listHeadtwo, listTail *Node
	var num int
	var flag bool
	pre = nil
	current = head
	for current != nil {
		num++
		if num == 1 {
			listHeadone = listHeadtwo
			listHeadtwo = current
		}
		next = current.Next
		current.Next = pre
		if num == k {
			if !flag {
				ret = current
				flag = true
			}else{
				listTail = current
				if listHeadone != nil {
					listHeadone.Next = listTail
				}
			}
			num = 0
			pre = nil
		}else{
			pre = current
		}
		current = next
	}
	if num != 0 {
		//listHeadone.Next = pre
		tmp := pre
		//反转回去最后不足k个节点
		var tmpPre, tmpNext *Node
		for tmp != nil {
			tmpNext = tmp.Next
			tmp.Next = tmpPre
			tmpPre = tmp
			tmp = tmpNext
		}
		if listHeadone != nil {
			listHeadone.Next = tmpPre
		}
	}else{
		if listHeadone != nil {
			listHeadone.Next = listTail
		}
	}
	if !flag {
		return head
	}
	return ret
}

func leetcode32(s string) int {
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
		//fmt.Println(tmplen, "---", simg, "---", k)
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
	l := len(nums)
	contains := 0
	for i := 0; i < l; i++ { //检查是否包含1
		if nums[i] == 1 {
			contains++
			break
		}
	}

	if contains == 0 { //不包含1即答案
		return 1
	}
	if l == 1 { //包含且长度为1，  2即答案
		return 2
	}

	for i := 0; i < l; i++ { //把所有大数,负数，全部转换为1，因为值必定 res <= l+1
		if nums[i] <= 0 || l < nums[i] {
			nums[i] = 1
		}
	}

	for i := 0; i < l; i++ { //以符号的正负 当作hash表
		index := AbsInt(nums[i]) //获取需要改变的索引
		if index == l {
			nums[0] = - AbsInt(nums[0])
		} else {
			nums[index] = - AbsInt(nums[index])
		}
	}

	for i := 1; i < l; i++ { //不存在 即答案
		if nums[i] > 0 {
			return i
		}
	}

	if nums[0] > 0 { // 不存在 即答案
		return l
	}
	return l + 1 //都存在，后序l+1为答案
}

func AbsInt(n int) int {
	if n > 0 {
		return n
	}
	return -n
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

//hai mei nong hao
func leetcode51(n int) [][]string {
	var ret [][]string
	empty := make([][]byte, n, n)
	for i := 0; i < n; i++ {
		empty[i] = make([]byte, n, n)
		for j := 0; j < n; j++ {
			empty[i][j] = '.'
		}
	}
	solveNQueens(n, &ret, empty, 0)
	//fmt.Println("chang is :", ret)
	return ret
}

func solveNQueens(n int, ret *[][]string, tmp [][]byte, hasPut int) {
	nextHasput := hasPut + 1
	if hasPut == n {
		tmpRet := make([]string, 0, n)
		for _,v := range tmp {
			tmpRet = append(tmpRet, string(v))
		}
		*ret = append(*ret, tmpRet)
		return
	}
	var flag bool
	for i := 0; i < n; i++ {
		flag = false
		tmp[hasPut][i] = 'Q'
		//fmt.Println(tmp, "--- i: ",i,"   hasput:  ",hasPut)
		//除去列
		for j := 0; j < hasPut; j++ {
			if tmp[j][i] == 'Q' {
				tmp[hasPut][i] = '.'
				flag = true
				break
			}
		}
		if flag {
			continue
		}
		//除去斜
		index := 1
		for {
			if i + index < n && hasPut + index < n {
				if tmp[hasPut+index][i+index] == 'Q' {
					tmp[hasPut][i] = '.'
					flag = true
					break
				}
			}else{
				break
			}
			index++
		}
		if flag {
			continue
		}
		index = 1
		for {
			if i - index >= 0 && hasPut - index >= 0 {
				if tmp[hasPut-index][i-index] == 'Q' {
					tmp[hasPut][i] = '.'
					flag = true
					break
				}
			}else{
				break
			}
			index++
		}
		if flag {
			continue
		}
		index = 1
		for {
			if i + index < n && hasPut - index >= 0 {
				if tmp[hasPut-index][i+index] == 'Q' {
					tmp[hasPut][i] = '.'
					flag = true
					break
				}
			}else{
				break
			}
			index++
		}
		if flag {
			continue
		}
		index = 1
		for {
			if i - index >= 0 && hasPut + index < n {
				if tmp[hasPut+index][i-index] == 'Q' {
					tmp[hasPut][i] = '.'
					flag = true
					break
				}
			}else{
				break
			}
			index++
		}
		if flag {
			continue
		}
		solveNQueens(n, ret, tmp, nextHasput)
		//还原原位
		tmp[hasPut][i] = '.'
	}
}

func leetcode55(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	var ret bool
	var maxIndex int
	for i := 0; i < len(nums) - 1; i++ {
		if i + nums[i] > maxIndex {
			maxIndex = i + nums[i]
			if maxIndex >= len(nums) - 1 {
				return true
			}
		}
		if i > maxIndex {
			return ret
		}
		if i == maxIndex && nums[i] == 0 {
			if i == len(nums) - 1 {
				return true
			}else{
				return false
			}
		}
	}
	return ret
}

type sort56 [][]int
func (o sort56) Len() int 				{ return len(o)}
func (o sort56) Less(i, j int) bool 	{ return o[i][0] < o[j][0] }
func (o sort56) Swap(i, j int)      	{ o[i], o[j] = o[j], o[i] }
func leetcode56(intervals [][]int) [][]int {
	if len(intervals) <=1 {
		return intervals
	}
	sort.Sort(sort56(intervals))
	ret := make([][]int, 0)
	tmpRegion := make([]int, 2, 2)
	for k, v := range intervals {
		if k == 0 {
			tmpRegion = v
			continue
		}
		if v[0] > tmpRegion[1] {
			ret = append(ret, tmpRegion)
			tmpRegion = v
			continue
		}else{
			if v[1] > tmpRegion[1] {
				tmpRegion[1] = v[1]
			}
		}
	}

	return append(ret, tmpRegion)
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

func leetcode65(s string) bool {
	//var ret, point, e, plusMinus bool

	//return ret
	return false
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

//两个字符的编辑距离
func leetcode72(word1, word2 string) int {
	len1 := len(word1)
	len2 := len(word2)
	dp := make([][]int, len1+1, len1+1)
	var min int
	for k,_ := range dp {
		if k == 0 {
			for x:=0; x<=len2; x++ {
				dp[k] = append(dp[k], x)
			}
		}else{
			dp[k] = make([]int, len2+1,len2+1)
			dp[k][0] = k
		}
	}
	for i:=1; i<len1+1; i++ {
		for j:=1; j<len2+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			}else{
				min = math.MaxInt64
				if j-1>=0 && dp[i][j-1] < min {
					min = dp[i][j-1]
				}
				if i-1>=0 && dp[i-1][j] < min {
					min = dp[i-1][j]
				}
				if i-1>=0 && j-1>=0 && dp[i-1][j-1] < min {
					min = dp[i-1][j-1]
				}
				dp[i][j] = min + 1
			}
		}
	}
	return dp[len1][len2]
}

//单调栈
func leetcode84(height []int) int {
	var ret int
	if height == nil || len(height) == 0 {
		return 0
	}
	ksline, vslice := make([]int, 0), make([]int, 0)
	for k, v := range height {
		if len(ksline) == 0 {
			ksline = append(ksline, k)
			vslice = append(vslice, v)
		} else {
			if v >= vslice[len(vslice)-1] {
				ksline = append(ksline, k)
				vslice = append(vslice, v)
			} else {
				i := len(vslice) - 1
				for {
					if i < 0 {
						vslice, ksline = make([]int, 0), make([]int, 0)
						ksline = append(ksline, k)
						vslice = append(vslice, v)
						//fmt.Println("append ", k, v)
						break
					}
					if vslice[i] < v {
						//fmt.Println("cut ", i, v, vslice, ksline)
						vslice = vslice[:i+1]
						vslice = append(vslice, v)
						ksline = ksline[:i+1]
						ksline = append(ksline, k)
						//fmt.Println("after cut", vslice, ksline)
						break
					}
					ans := (k - ksline[i])*vslice[i]
					if ans > ret {
						ret = ans
					}
					i--
				}
			}
		}
		//fmt.Println(vslice, ksline)
	}
	//fmt.Println("final ", vslice, ksline)
	i := len(vslice) - 1
	for j:=0; j<i; j ++{
		ans := (ksline[i] - ksline[j])*vslice[j]
		if ans > ret {
			ret = ans
		}
	}
	return ret
}

func leetcode85(matrix [][]byte) int {
	var ret int
	if len(matrix) != 0 {
		height := make([]int, len(matrix[0]), len(matrix[0]))
		for row:=0; row<len(matrix); row++ {
			for col:=0; col<len(matrix[0]); col++ {
				if matrix[row][col] == '1' {
					height[col]++
				} else {
					height[col] = 0
				}
			}
			ans := leetcode84(height)
			if ans > ret {
				ret = ans
			}
		}
	}
	return ret
}

func leetcode91(s string) int {
	var ret int
	sb := []byte(s)
	l := len(sb)
	if l == 0 || sb[0] == '0' {
		return 0
	}
	dp := make([]int, l, l)
	dp[0] = 1
	dp[1] = 1
	for i:=1; i<l; i++ {

	}
	return ret
}

func leetcode93(s string) []string {
	var ret []string
	l := len(s)
	if l < 4 || l > 12 {
		return ret
	}
	dfs93([]byte(s), make([]byte, 0), 0, 0, &ret)
	return ret
}
func dfs93(s []byte, ans []byte, count, index int, ret *[]string) {
	if count >= 4 {
		if len(s) == index {
			*ret = append(*ret, string(ans))
		}
		return
	}
	if index + 1 <= len(s) {
		tmp := make([]byte, len(ans))
		copy(tmp, ans)
		tmp = append(tmp, s[index])
		if count + 1 < 4 {
			tmp = append(tmp, '.')
		}
		dfs93(s, tmp, count+1, index+1, ret)
	}
	if index + 2 <= len(s) {
		tmp := make([]byte, len(ans))
		copy(tmp, ans)
		if s[index] == '0' {
			return
		}
		tmp = append(tmp, s[index:index+2]...)
		if count + 1 < 4 {
			tmp = append(tmp, '.')
		}
		dfs93(s, tmp, count+1, index+2, ret)
	}
	if index + 3 <= len(s) {
		tmp := make([]byte, len(ans))
		copy(tmp, ans)
		if s[index] == '0' || s[index] > '2'{
			return
		}
		if s[index] == '2' {
			if s[index+1] > '5' {
				return
			} else if s[index+1] == '5' {
				if s[index+2] > '5' {
					return
				}
			}
		}
		tmp = append(tmp, s[index:index+3]...)
		if count + 1 < 4 {
			tmp = append(tmp, '.')
		}
		dfs93(s, tmp, count+1, index+3, ret)
	}
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