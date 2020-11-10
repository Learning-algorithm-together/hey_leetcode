package leetcode

/**
请实现一个函数用来匹配包含'. '和'*'的正则表达式。模式中的字符'.'表示任意一个字符，
而'*'表示它前面的字符可以出现任意次（含0次）。在本题中，匹配是指字符串的所有字符匹配整个模式。
例如，字符串"aaa"与模式"a.a"和"ab*ac*a"匹配，但与"aa.a"和"ab*a"均不匹配。

示例 1:

输入:
s = "aa"
p = "a"
输出: false
解释: "a" 无法匹配 "aa" 整个字符串。
示例 2:

输入:
s = "aa"
p = "a*"
输出: true
解释:因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
示例3:

输入:
s = "ab"
p = ".*"
输出: true
解释:".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。
示例 4:

输入:
s = "aab"
p = "c*a*b"
输出: true
解释:因为 '*' 表示零个或多个，这里 'c' 为 0 个, 'a' 被重复一次。因此可以匹配字符串 "aab"。
示例 5:

输入:
s = "mississippi"
p = "mis*is*p*."
输出: false
s可能为空，且只包含从a-z的小写字母。
p可能为空，且只包含从a-z的小写字母以及字符.和*，无连续的 '*'。

*/
/**
递归
终止条件：p为空时，s为空返回true，s不空返回false
first标志首位是否相等，p[0]是.也表示相等
(1)如果p的第2位为*
①首位相同 则等同于判断s[1:]和p （忽略s[0]）
②首位不同 则等同于判断s和p[2:] （将p前两位x*和空值配对）
(2)否则不要期待奇迹出现了，
首位如果相等就各向后移一位，判断s[1:]和p[1:]
如果不等就是false
这里利用了go字符串的截断符号

*/
func isMatch2(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	var first bool
	if len(s) > 0 && (s[0] == p[0] || p[0] == '.') {
		first = true
	}
	if len(p) > 1 && p[1] == '*' {
		return (first && isMatch(s[1:], p)) || isMatch(s, p[2:])
	}
	return first && isMatch(s[1:], p[1:])
}

func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	var first bool
	if len(s) > 0 && (s[0] == p[0] || p[0] == '.') {
		first = true
	}
	if len(p) > 1 && p[1] == '*' {
		//if first {
		//	return isMatch(s[1:], p)
		//} else {
		//	return isMatch(s, p[2:])
		//}
		return (first && isMatch(s[1:], p)) || isMatch(s, p[2:]) //相当于上面注释的写法。这种写法很高级

	}
	return first && isMatch(s[1:], p[1:])
}
