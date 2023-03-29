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