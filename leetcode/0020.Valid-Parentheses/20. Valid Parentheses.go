package leetcode

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	stack := make([]rune, 0)
	for _, v := range s {
		size := len(stack)
		if v == '(' {
			stack = append(stack, ')')
		} else if v == '[' {
			stack = append(stack, ']')
		} else if v == '{' {
			stack = append(stack, '}')
		} else if size == 0 || stack[size-1] != v {
			return false
		} else {
			stack = stack[:size-1]
		}
	}
	return len(stack) == 0
}
