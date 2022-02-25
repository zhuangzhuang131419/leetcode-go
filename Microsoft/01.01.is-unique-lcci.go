package leetcode

func isUnique(astr string) bool {
	var mark int32
	for i := 0; i < len(astr); i++ {
		if mark&(1<<(astr[i]-'a')) != 0 {
			return false
		} else {
			mark = mark | (1 << (astr[i] - 'a'))
		}

	}
	return true
}
