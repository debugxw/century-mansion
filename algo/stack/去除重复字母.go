package stack

import "strings"

// 力扣 316、去除重复字母
// 给你一个字符串 s ，请你去除字符串中重复的字母，使得每个字母只出现一次
// 需保证 返回结果的字典序最小（要求不能打乱其他字符的相对位置）
func removeDuplicateLetters(s string) string {
	cntMap := map[int32]int{}
	for _, ch := range s {
		cntMap[ch]++
	}
	inStack := map[int32]bool{}
	var stack []int32
	for _, char := range s {
		cntMap[char]--
		if _, exist := inStack[char]; exist && inStack[char] {
			continue
		}
		for len(stack) != 0 && char < stack[len(stack)-1] && cntMap[stack[len(stack)-1]] > 0 {
			inStack[stack[len(stack)-1]] = false
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, char)
		inStack[char] = true
	}
	ret := strings.Builder{}
	for i := range stack {
		ret.WriteString(string(stack[i]))
	}
	return ret.String()
}
