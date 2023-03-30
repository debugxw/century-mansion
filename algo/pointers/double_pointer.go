package pointers

// LengthOfLongestSubstring 力扣 3、无重复字符的最长子串
// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度
func LengthOfLongestSubstring(s string) int {
	m := map[byte]struct{}{}
	resL, resR := -1, -1
	l, r := 0, 0
	for r < len(s) {
		for l <= r {
			if _, exist := m[s[r]]; exist {
				delete(m, s[l])
				l++
				continue
			}
			break
		}
		m[s[r]] = struct{}{}
		r++
		if r-l > resR-resL {
			resL, resR = l, r
		}
	}
	return resR - resL
}

// MinWindow 力扣 76、最小覆盖子串
// 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 ""
func MinWindow(s string, t string) string {
	tMap := map[byte]int{}
	for _, c := range t {
		tMap[byte(c)]++
	}
	sMap := map[byte]int{}
	l, r := 0, 0
	min := s
	for r < len(s) {
		sMap[s[r]]++
		r++
		for l < r && check(sMap, tMap) {
			if len(min) > r-l {
				min = s[l:r]
			}
			sMap[s[l]]--
			l++
		}
	}
	if l == 0 && r == len(s) {
		return ""
	}
	return min
}

func check(s, t map[byte]int) bool {
	for key, val := range t {
		if _, exist := s[key]; !exist || s[key] < val {
			return false
		}
	}
	return true
}

// FindAnagrams 438、找到字符串中所有字母异位词
// 给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词的子串，返回这些子串的起始索引。不考虑答案输出的顺序
// 异位词 指由相同字母重排列形成的字符串（包括相同的字符串）
func FindAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return nil
	}
	pMap, missMap := map[byte]struct{}{}, map[byte]int{}
	for _, c := range p {
		pMap[byte(c)] = struct{}{}
		missMap[byte(c)]++
	}
	var ret []int
	pLen := len(p)
	i := 0
	for i < len(s) {
		if i >= pLen {
			if _, exist := pMap[s[i-pLen]]; exist {
				missMap[s[i-pLen]]++
				if missMap[s[i-pLen]] == 0 {
					delete(missMap, s[i-pLen])
				}
			}
		}
		if _, exist := pMap[s[i]]; exist {
			missMap[s[i]]--
			if missMap[s[i]] == 0 {
				delete(missMap, s[i])
			}
		}
		if i >= pLen-1 && len(missMap) == 0 {
			ret = append(ret, i-pLen+1)
		}
		i++
	}
	return ret
}

// BalancedString 力扣 1234、替换子串得到平衡字符串
// 有一个只含有'Q', 'W', 'E', 'R'四种字符，且长度为 n 的字符串，n是4的倍数
// 假如在该字符串中，这四个字符都恰好出现n/4次，那么它就是一个「平衡字符串」
// 给你一个这样的字符串 s，请通过「替换一个子串」的方式，使原字符串 s 变成一个「平衡字符串」。
// 你可以用和「待替换子串」长度相同的任何其他字符串来完成替换。
// 请返回待替换子串的最小可能长度
func BalancedString(s string) int {
	m := map[byte]int{}
	for _, val := range s {
		m[byte(val)]++
	}
	avgCnt := len(s) / 4
	if checkBalance(&m, avgCnt) {
		return 0
	}

	l, r, min := 0, 0, len(s)
	for r < len(s) {
		m[s[r]]--
		r++
		for l < r && checkBalance(&m, avgCnt) {
			if r-l < min {
				min = r - l
			}
			m[s[l]]++
			l++
		}
	}
	return min
}

func checkBalance(m *map[byte]int, avg int) bool {
	for _, val := range *m {
		if val > avg {
			return false
		}
	}
	return true
}